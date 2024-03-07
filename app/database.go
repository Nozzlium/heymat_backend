package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func initDB() (*sql.DB, error) {
	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"heymat",
		"yukberhemat",
		"heymat",
	)
	return sql.Open("postgres", url)
}
