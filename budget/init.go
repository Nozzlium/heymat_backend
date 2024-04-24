package budget

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	DB             *sql.DB
	AuthMiddleware func(*fiber.Ctx) error
}

// func Init(
// 	app *fiber.App,
// 	config Config,
// ) error {
// 	if config.db == nil {
// 		return errors.new(
// 			"cannot init, db is nil",
// 		)
// 	}
// 	db = config.db
// 	budget := app.Group("/api/budget")
// 	budget.Use(config.AuthMiddleware)
// 	budget.Post("", createBudgetHandler)
// 	budget.Get(
// 		"",
// 		getBudgetPlanItemsHandler,
// 	)
// 	budget.Get(
// 		"/:budgetPlanId",
// 		getBudgetPlanByIdHandler,
// 	)
// 	if config.AuthMiddleware == nil {
// 		return errors.New(
// 			"cannot init, auth middleware is nil",
// 		)
// 	}
// 	return nil
// }

func Init(
	config Config,
) (*fiber.App, error) {
	if config.DB == nil {
		return nil, errors.New(
			"cannot init, db is nil",
		)
	}
	DB = config.DB
	budget := fiber.New()
	if config.AuthMiddleware == nil {
		return nil, errors.New(
			"auth middleware is nil, cannot init budget route",
		)
	}
	budget.Use(config.AuthMiddleware)

	budget.Post(
		"",
		createBudgetHandler,
	)
	budget.Get(
		"",
		getBudgetPlanItemsHandler,
	)
	budget.Get(
		"/:budgetPlanId",
		getBudgetPlanByIdHandler,
	)
	return budget, nil
}
