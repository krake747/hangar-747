package core

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Run init script
	initSQL, err := os.ReadFile("testdata/init.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(initSQL))
	if err != nil {
		return nil, err
	}

	return db, nil
}
