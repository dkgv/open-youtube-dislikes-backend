package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("environment variable DATABASE_URL is not set")
	}

	prefix := "?"
	if strings.Contains(databaseUrl, "?") {
		prefix = "&"
	}
	databaseUrl += prefix + "options=-csearch_path%3Dopen_youtube_dislikes"

	log.Println("Connecting to database:", databaseUrl)
	conn, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	err = migrateSchemas(conn)
	if err != nil {
		return nil, err
	}

	log.Println("Established connection to database")
	return conn, nil
}
