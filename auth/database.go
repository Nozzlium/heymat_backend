package auth

import (
	"database/sql"
	"errors"
)

var DB *sql.DB

func SetDatabaseInstance(db *sql.DB) {
	if db == nil {
		panic(
			errors.New(
				"cannot set database instance",
			),
		)
	}

	DB = db
}
