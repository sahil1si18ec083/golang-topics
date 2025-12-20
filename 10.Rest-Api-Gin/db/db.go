package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func INITDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("could not connect to Databse")
	}
	err = DB.Ping()
	if err != nil {
		panic("could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()

}
func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS Event (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME DEFAULT CURRENT_TIMESTAMP  NOT NULL,
	UserId INTEGER
	)
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
	}

}
