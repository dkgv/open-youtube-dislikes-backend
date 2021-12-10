package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

//go:embed internal/database/migrations/*.sql
var embeddedFiles embed.FS

func migrateSchemas(conn *sql.DB) error {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("connecting to database instance failed: %w", err)
	}

	migrationSource, err := NewEmbeddedFileSource("internal/database/migrations")
	if err != nil {
		return fmt.Errorf("database migration failed: %w", err)
	}
	migration, err := migrate.NewWithInstance("", migrationSource, "postgres", driver)
	if err != nil {
		return fmt.Errorf("database migration with instance failed: %w", err)
	}

	if err := migration.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("database migration failed: %w", err)
		}
	}

	return nil
}
