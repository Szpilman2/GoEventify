package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //we add _ to prevent removing this package. because we don't use it directly and we use database/sql instead.
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") //open a connection to database.

	if err != nil {
		panic("could not connect to database.")
	}
	DB.SetMaxOpenConns(10)

	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name  TEXT NOT NULL,
			description TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("could not create events table %v", err)
	}
}
