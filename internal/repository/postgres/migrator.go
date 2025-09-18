package postgres

import (
	"database/sql"
	"fmt"

		"github.com/golang-migrate/migrate/v4"
		"github.com/golang-migrate/migrate/v4/database/postgres"
		_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(conn *sql.DB) {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("migration driver could not be created: %v", err))
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/repository/migrations/",
		"postgres",
		driver,
	)
	if err != nil {
		panic(fmt.Sprintf("couldn't create a migrator: %v", err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Sprintf("migrations could not be applied: %v", err))
	}
}
