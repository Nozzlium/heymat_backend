package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/nozzlium/heymat_backend/lib"
)

//go:embed migrations/*.sql
var schema embed.FS

func InitDB(
	config *lib.Config,
) (*sql.DB, *migrate.Migrate) {
	url := config.GetPGConnString()
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	driver, err := postgres.WithInstance(
		db,
		&postgres.Config{},
	)
	if err != nil {
		panic(err)
	}

	sourceDriver, err := iofs.New(
		schema,
		"migrations",
	)
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		sourceDriver,
		"heymat",
		driver,
	)
	if err != nil {
		panic(err)
	}

	return db, m
}

func Migrate(
	m *migrate.Migrate,
	steps uint8,
) {
	if steps == 0 {
		err := m.Up()
		if err != nil {
			if err == migrate.ErrNoChange {
				return
			}
			panic(err)
		}
		return
	}

	err := m.Steps(int(steps))
	if err != nil {
		panic(err)
	}
}

func Truncate(m *migrate.Migrate) {
	err := m.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		panic(err)
	}
}

func GetDatabaseVersion(
	m *migrate.Migrate,
) {
	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			fmt.Println(
				"no database verion, migrate with --migrate-up command",
			)
			return
		}
		log.Println(err)
		panic(err)
	}

	if dirty {
		fmt.Printf(
			"dirty state at version %d\n",
			version,
		)
		return
	}

	fmt.Printf(
		"current migration version: %d\n",
		version,
	)
}
