package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nozzlium/heymat_backend/auth"
	"github.com/nozzlium/heymat_backend/budget"
	"github.com/nozzlium/heymat_backend/lib"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return lib.WriteErrorResponse(
				fiber.StatusInternalServerError,
				err,
				ctx,
			)
		},
	})
	app.Use(cors.New())

	auth.GetRouting(app)
	budget.GetRouting(app)

	err := app.Listen(":4343")
	if err != nil {
		panic(err)
	}
}
