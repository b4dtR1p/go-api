Simple REST Api in go.

try to install directy with `go install github.com/b4dtR1p/go-api`

and then start server with a simple ./go-api inside repo dir or try to install on your system with go install from repo dir;

Open your browser and go to `localhost:8080/api/items` or simple CUrl from cli;

Create new Items with a simple `curl -H "Content-Type: application/json" -d '{"Name": "iPhone-6S", "Picture": "http://cdn0.vox-cdn.com/uploads/chorus_asset/file/798874/DSCF1913.0.jpg", "Description": "iPhone 6S - 64Gb Withe-Gold", "Price": "$890,00"}' localhost:8080/api/items`

Have Fun Hackers!
