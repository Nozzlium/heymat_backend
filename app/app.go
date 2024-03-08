package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/controllers"
	"github.com/nozzlium/heymat_backend/custom_errors"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/services"
)

func InitApp() (*fiber.App, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	userRepository := repositories.NewUserRepository()

	authService := services.NewAuthService(userRepository, db)
	userService := services.NewUserService(userRepository, db)

	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)

	app := fiber.New(fiber.Config{
		ErrorHandler: custom_errors.ErrorHandler,
	})
	authRouting(app, authController)
	userRouting(app, userController)

	return app, nil
}
