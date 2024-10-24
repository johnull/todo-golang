package database

import (
	"database/sql"
	"log"
)

func setupDatabase(db *sql.DB) error {
	if _, err := db.Exec(`CREATE DATABASE IF NOT EXISTS todoDB`); err != nil {
		log.Printf("database error: %v", err)
		return err
	}

	if _, err := db.Exec(`USE todoDB`); err != nil {
		log.Printf("database error: %v", err)
		return err
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS TodoList (
			id INT AUTO_INCREMENT,
			item TEXT NOT NULL,
			completed SMALLINT DEFAULT 0,
			PRIMARY KEY (id)
		);
	`); err != nil {
		log.Printf("database error: %v", err)
		return err
	}

	return nil
}
