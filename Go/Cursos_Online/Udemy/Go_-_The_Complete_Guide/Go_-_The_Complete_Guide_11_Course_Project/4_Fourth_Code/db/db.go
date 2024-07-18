package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // We use the "_" to specify that we will be using the library indirectly (It will help enable functionality to a third library, in this case "sql")
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create event table.")
	}
}
