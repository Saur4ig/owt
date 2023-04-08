package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableUsers = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		address TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		phone TEXT NOT NULL UNIQUE
	);`

	tableSkills = `
	CREATE TABLE IF NOT EXISTS skills (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
	);`

	tableUsersSkills = `
	CREATE TABLE IF NOT EXISTS user_skills (
		user_id INTEGER,
		skill_id INTEGER ,
		level INTEGER NOT NULL,
		FOREIGN KEY (user_id)  REFERENCES users (id)
		FOREIGN KEY (skill_id)  REFERENCES skills (id)
		primary key (user_id, skill_id)
	);`
)

func Init(dbName string) (*sql.DB, error) {
	db, err := openWithInit(dbName)
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func openWithInit(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s?parseTime=true", dbName))
	if err != nil {
		return nil, err
	}

	// maximum number of open connections
	db.SetMaxOpenConns(10)

	return db, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(tableUsers)
	if err != nil {
		return err
	}

	_, err = db.Exec(tableSkills)
	if err != nil {
		return err
	}

	_, err = db.Exec(tableUsersSkills)
	if err != nil {
		return err
	}

	return nil
}
