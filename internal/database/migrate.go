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

	migrationSource, err := NewEmbeddedFileSource("migrations")
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("", migrationSource, "postgres", driver)
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

func NewEmbeddedFileSource(path string) (source.Driver, error) {
	embeddedSource := &EmbbedSource{fs: embeddedFiles}
	sourceDriver, err := embeddedSource.Open(path)
	if err != nil {
		return nil, err
	}

	return sourceDriver, nil
}

func init() {
	source.Register("embed", &EmbbedSource{})
}

func (e *EmbbedSource) Open(path string) (source.Driver, error) {
	es := &EmbbedSource{fs: e.fs}
	if err := es.Init(http.FS(es.fs), path); err != nil {
		return nil, err
	}

	return es, nil
}
