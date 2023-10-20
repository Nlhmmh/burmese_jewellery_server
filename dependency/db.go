package dependency

import (
	"database/sql"
	"time"
)

var postgresDB *sql.DB

func newDB(dbSource string) error {
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	postgresDB = db

	return nil
}
