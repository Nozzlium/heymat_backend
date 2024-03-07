package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/controllers"
)

func userRouting(app *fiber.App, userController controllers.UserController) {
	app.Post("/register", userController.Register)
}
