package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nozzlium/heymat_backend/custom_errors"
	"github.com/nozzlium/heymat_backend/routing"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: custom_errors.ErrorHandler,
	})
	app.Use(cors.New())
	routing.Routing(app)

	err := app.Listen(":4343")
	if err != nil {
		panic(err)
	}
}
