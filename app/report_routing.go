package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/controllers"
	"github.com/nozzlium/heymat_backend/middleware"
)

func reportRouting(app *fiber.App, reportController controllers.ReportController) {
	balance := app.Group("/balance")
	balance.Use(middleware.AuthMiddleware)
	balance.Get("", reportController.GetByYear)
}
