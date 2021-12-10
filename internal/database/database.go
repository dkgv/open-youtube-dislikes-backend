package database

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("environment variable DATABASE_URL is not set")
	}

	databaseUrl += "?options=-csearch_path%3Ddislikes"

	conn, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	err = migrateSchemas(conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
