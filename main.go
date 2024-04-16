package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nozzlium/heymat_backend/auth"
	"github.com/nozzlium/heymat_backend/budget"
	"github.com/nozzlium/heymat_backend/lib"
)

func main() {
	lib.LoadEnv()
	config := lib.InitConfig()

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

	lib.SetDatabaseInstance(db)
	budget.SetDatabaseInstance(db)
	auth.SetDatabaseInstance(db)

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

	err := app.Listen(":4040")
	if err != nil {
		panic(err)
	}
}
