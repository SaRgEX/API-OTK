package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

func MigrateSQL(conn *sqlx.DB, driverName string) error {
	driver, _ := postgres.WithInstance(conn.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		driverName,
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
