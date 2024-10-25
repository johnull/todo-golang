# Todo Golang
Basic to-do list application

### Installation
1. Create and modify `.env`
```bash
DB_HOST=
DB_USER=
DB_PASS=
```
2. Install dependencies
```bash
$ go get ./...
```

### Local run
Setup database, run mysql/mariadb
```bash
source <complete_path>/db/migration/000001_create_items_table.up.sql
```
Run app
```
$ go run cmd/webapp/main.go

http://localhost:8080
```
### Docker run
```bash
$ docker compose up -d
```
Setup database
```bash
$ docker run -v <complete_path>/db/migration:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://root:root@tcp(localhost:3306)/todoDB" up
```
Run app
```bash
http://localhost:8080
```
### Structure
```
├── cmd
│   └── webapp
│       └── main.go
├── controllers
│   └── todo.go
├── database
│   ├── connection.go
│   └── setup_db.go
├── go.mod
├── go.sum
├── models
│   ├── todo.go
│   └── todo_view.go
└── views
    └── index.html
```
