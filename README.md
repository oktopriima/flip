# flip

## installation
This application required [Golang](https://golang.org/) v 1.15+

Install the dependencies and devDependencies and start the server.

```sh
go get -u github.com/oktopriima/flip
cd $GOPATH/src/github.com/oktopriima/flip/
go mod tidy
go mod vendor

cp env-example.yaml env.yaml
cp dbconfig-example.yml dbconfig.yml

# library to sql migrate
# helping for make a database migration
go get -v github.com/rubenv/sql-migrate/...

# adjust env.yaml and dbconfig.yml
sql-migrate up --env=local

go build
```
after  `go build` command execute, there's an executeable file.
you can just run the executeable file, and then open your browser and access [http://localhost:5000](http://localhost:5000)
