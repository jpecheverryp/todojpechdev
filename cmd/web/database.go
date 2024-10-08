package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
