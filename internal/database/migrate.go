package database

import (
	"database/sql"
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
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

	embeddedSource := &EmbbedSource{fs: embeddedFiles}
	sourceDriver, err := embeddedSource.Open("migrations")
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

type EmbbedSource struct {
	httpfs.PartialDriver
	fs embed.FS
}

func (e *EmbbedSource) Open(path string) (source.Driver, error) {
	es := &EmbbedSource{fs: e.fs}
	if err := es.Init(http.FS(es.fs), path); err != nil {
		return nil, err
	}

	return es, nil
}

func init() {
	source.Register("embed", &EmbbedSource{})
}
