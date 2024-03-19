package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/controllers"
	"github.com/nozzlium/heymat_backend/middleware"
)

func budgetRouting(app *fiber.App, budgetController controllers.BudgetController) {
	budget := app.Group("/budget")
	budget.Use(middleware.AuthMiddleware)
	budget.Post("", budgetController.Create)
}
