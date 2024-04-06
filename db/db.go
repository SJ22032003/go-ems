package db

import (
	sql "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var db_err error

func InitDB() {
	DB, db_err = sql.Open("sqlite3", "app.db")
	if db_err != nil {
		panic(db_err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func Execute(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Query(args...)

}

func createTables() {

	var err error

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err = Execute(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	 )
`
	_, err = Execute(createEventsTable)
	if err != nil {
		panic(err)
	}

}
