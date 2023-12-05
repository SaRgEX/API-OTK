package migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func MigrateSQL(conn *sqlx.DB, driverName string) error {
	m, err := migrate.New(
		"file://db//migrations",
		"postgres://postgres:123@localhost:5436/postgres?sslmode=disable",
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Migrations applied")

	return nil
}
