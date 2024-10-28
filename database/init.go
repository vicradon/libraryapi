package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		genre TEXT,
		is_available BOOLEAN,
		edition TEXT,
		summary TEXT
	)`)

	if err != nil {
		log.Fatal(err)
	}
}
