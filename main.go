package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nozzlium/heymat_backend/auth"
	"github.com/nozzlium/heymat_backend/budget"
	"github.com/nozzlium/heymat_backend/lib"
)

func main() {
	LoadEnv()
	config := InitConfig()

	lib.SetSignKey()
	db, m := InitDB(config)

	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "--migrate-up":
			Migrate(m, 0)
			break
		case "--version":
			GetDatabaseVersion(m)
			return
		case "--truncate":
			Truncate(m)
			return
		}
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return lib.WriteResponse(
				ctx,
				fiber.StatusInternalServerError,
				err.Error(),
			)
		},
	})
	app.Use(cors.New())

	authRoute, err := auth.Init(
		auth.Config{
			DB: db,
		},
	)
	if err != nil {
		log.Fatal(
			"Initializing auth route",
			err.Error(),
		)
	}
	app.Mount("", authRoute)

	budgetRoute, err := budget.Init(
		budget.Config{
			DB:             db,
			AuthMiddleware: auth.AuthMiddleware,
		},
	)
	if err != nil {
		log.Fatal(
			"Initializing budget route",
			err.Error(),
		)
	}
	app.Mount(
		"/api/budget",
		budgetRoute,
	)

	err = app.Listen(":4040")
	if err != nil {
		log.Fatal(err.Error())
	}
}
