package database

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	if err := godotenv.Load("../../.env"); err != nil {
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
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/", username, password, host)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	
	if err := setupDatabase(db); err != nil {
		log.Fatalf("setup database error: %v", err)
	}
	
	log.Println("Connected to database")
	
	return db
}
