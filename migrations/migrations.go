package migrations

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	migrationsSchemaName = "public"
	driverName           = "pgx"
	pathToMigrations     = "file://migrations/"
)

func Run(db *sql.DB) (uint, error) {
	query := `create schema if not exists ` + migrationsSchemaName
	if _, err := db.Exec(query); err != nil {
		return 0, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{SchemaName: migrationsSchemaName})
	m, err := migrate.NewWithDatabaseInstance(pathToMigrations, driverName, driver)
	if err != nil {
		return 0, err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return 0, err
	}
	version, _, _ := m.Version()
	return version, nil
}
