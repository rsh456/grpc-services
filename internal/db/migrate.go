package db

import (
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	// Run our migrations located within the migrations directory
	// These migrations will create our tables if they don't exist
	// and perfomr actions such as adding new columns.
	// These will run in order and basically acs as a list of steps
	// that eventually setup the database.

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("no change made by migrations")
		} else {
			return err
		}

	}
	return nil
}
