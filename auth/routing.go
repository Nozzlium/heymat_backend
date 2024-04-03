package auth

import "github.com/gofiber/fiber/v2"

func GetRouting(app *fiber.App) {
	app.Post("/api/login", loginHandler)
	app.Post(
		"/api/register",
		registerHandler,
	)
}
