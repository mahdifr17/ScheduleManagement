package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/app.db")
	if err != nil {
		log.Fatal(err)
	}

	dbMigration(db)

	return db
}

func dbMigration(db *sql.DB) {
	migrationStmnt := `
		CREATE TABLE IF NOT EXISTS users (
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS shift_schedule (
			id TEXT NOT NULL PRIMARY KEY
		);
	`
	_, err := db.Exec(migrationStmnt)
	if err != nil {
		log.Fatal(err)
	}
}
