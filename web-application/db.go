package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func ConnectToDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", name)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
