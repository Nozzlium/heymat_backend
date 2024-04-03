package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB = initDB()

func initDB() *sql.DB {
	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"heymat",
		"yukberhemat",
		"heymat",
	)
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	return db
}
