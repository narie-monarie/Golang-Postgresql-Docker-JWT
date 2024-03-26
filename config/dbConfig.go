package config

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
