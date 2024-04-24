package auth

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	DB *sql.DB
}

// func Init(
// 	app *fiber.App,
// 	config Config,
// ) error {
// 	if config.DB == nil {
// 		return errors.New(
// 			"db is nil, cannot set db",
// 		)
// 	}
// 	DB = config.DB
// 	auth := app.Group("")
// 	auth.Post(
// 		"/api/login",
// 		loginHandler,
// 	)
// 	auth.Post(
// 		"/api/register",
// 		registerHandler,
// 	)
// 	return nil
// }

func Init(
	config Config,
) (*fiber.App, error) {
	if config.DB == nil {
		return nil, errors.New(
			"db is nil, cannot init auth route",
		)
	}
	DB = config.DB
	auth := fiber.New()
	auth.Post(
		"/api/login",
		loginHandler,
	)
	auth.Post(
		"/api/register",
		registerHandler,
	)
	return auth, nil
}
