# Todo Golang
Basic to-do list application

#### Installation
1. Create and modify `.env`
```bash
DB_HOST=
DB_USER=
DB_PASS=
```
2. Install dependencies
```bash
go get -u github.com/joho/godotenv
go get -u github.com/go-sql-driver/mysql
```
3. Run
```
go run cmd/webapp/main.go

http://localhost:8080
```

#### Structure
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
