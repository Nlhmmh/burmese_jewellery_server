package server

import (
	"database/sql"
	"time"
)

// newDB - open database
func newDB(dbSource string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil

}
