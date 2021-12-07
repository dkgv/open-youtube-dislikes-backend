package database

import (
	"database/sql"
	"errors"
	"os"
)

func NewConnection() (*sql.DB, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("environment variable DATABASE_URL is not set")
	}

	return sql.Open("postgres", databaseUrl)
}
