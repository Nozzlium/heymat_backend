package budget

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/lib"
)

func GetRouting(app *fiber.App) {
	budget := app.Group("/api/budget")
	budget.Use(lib.AuthMiddleware)
	budget.Post("", createBudgetHandler)
	budget.Get(
		"",
		getBudgetPlanItemsHandler,
	)
	budget.Get(
		"/:budgetPlanId",
		getBudgetPlanByIdHandler,
	)
}
