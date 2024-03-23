package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/lib"
	"github.com/nozzlium/heymat_backend/middleware"
)

func Routing(instance *fiber.App) {
	authMiddleware := middleware.AuthMiddleware

	instance.Post("/api/login", lib.AuthController.Login)
	instance.Post("/api/register", lib.UserController.Register)

	budget := instance.Group("/budget")
	budget.Use(authMiddleware)
	budget.Post("", lib.BudgetController.Create)
	budget.Get("", lib.BudgetController.Get)
}
