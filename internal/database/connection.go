package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
var (
	_, b, _, _ = runtime.Caller(0)
	root = filepath.Join(filepath.Dir(b), "../../.env")
)
func ConnectDB() *sql.DB {
	if err := godotenv.Load(root); err != nil {
		log.Fatal(err)
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("DB_HOST not set")
	}
	username, ok := os.LookupEnv("DB_USER")
	if !ok {
		log.Fatal("DB_USER not set")
	}
	password, ok := os.LookupEnv("DB_PASS")
	if !ok {
		log.Fatal("DB_PASS")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/todoDB", username, password, host)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	log.Println("Connected to database")

	return db
}
