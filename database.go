package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"os"
	"path"
	"strconv"
	"time"
)

type Database struct {
	// Bolt database.
	db *bolt.DB
}

// Errors
var (
	ErrBucketNotFound = errors.New("bucket not found")
)

// Create and open a database.
func NewDatabase(filename string) (*Database, error) {
	// Open database
	os.MkdirAll(path.Dir(filename), 0700)
	db, err := bolt.Open(filename, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return nil, err
	}
	d := &Database{db}

	// Return
	return d, nil
}

// Close database
func (db *Database) Close() {
	db.db.Close()
	db.db = nil
}

func (db *Database) NewItemId() uint64 {
	var id uint64

	err := db.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("item"))
		if err != nil {
			return err
		}

		id, err = bucket.NextSequence()
		return err
	})

	if err != nil {
		panic(err)
	}

	return id
}

func (db *Database) SaveItem(item *Item) error {
	encoded, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return db.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("item"))
		if err != nil {
			return err
		}

		id := strconv.FormatUint(item.Id, 10)
		err = bucket.Put([]byte(id), encoded)
		if err != nil {
			return err
		}

		return nil
	})
}

func (db *Database) Item(id uint64) *Item {
	var item *Item = nil
	db.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("item"))
		if bucket == nil {
			return nil
		}

		sid := strconv.FormatUint(id, 10)

		c := bucket.Cursor()
		for k, v := c.Seek([]byte(sid)); bytes.Equal(k, []byte(sid)); k, v = c.Next() {
			json.Unmarshal(v, &item)
			return nil
		}
		return nil
	})
	return item
}

func (db *Database) Items() []*Item {
	var list []*Item
	db.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("item"))
		if bucket == nil {
			return nil
		}

		bucket.ForEach(func(k, v []byte) error {
			item := &Item{}
			json.Unmarshal(v, &item)
			list = append(list, item)
			return nil
		})
		return nil
	})
	return list
}

func (db *Database) RemoveItem(item *Item) error {
	return db.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("item"))
		if bucket == nil {
			return nil
		}

		id := strconv.FormatUint(item.Id, 10)
		return bucket.Delete([]byte(id))
	})
}
