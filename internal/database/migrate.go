package database

import (
	"database/sql"
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var embeddedFiles embed.FS

func migrateSchemas(conn *sql.DB) error {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return err
	}

	sourceDriver, err := httpfs.New(http.FS(embeddedFiles), "migrations")
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("", sourceDriver, "postgres", driver)
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
