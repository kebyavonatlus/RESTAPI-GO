package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (database *Database) MigrateDB() error {
	fmt.Println("migrating our database")

	driver, err := postgres.WithInstance(database.Client.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := migration.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migration: %w", err)
		}
	}

	fmt.Println("successfully migrated the database")

	return nil
}
