package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/controllers"
)

func authRouting(app *fiber.App, authController controllers.AuthController) {
	app.Post("/login", authController.Login)
}
