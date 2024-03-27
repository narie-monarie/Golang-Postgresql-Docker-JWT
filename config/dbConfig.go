package config

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	db, err := sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
