package lib

import (
	"database/sql"
	"errors"
)

var DB *sql.DB

func SetDatabaseInstance(db *sql.DB) {
	if DB == nil {
		panic(
			errors.New(
				"cannot set DB instance",
			),
		)
	}

	DB = db
}
