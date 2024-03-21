package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	budgetRepository := repositories.NewBudgetRepository()
	recordEntryRepository := repositories.NewRecordEntryRepository()

	authService := services.NewAuthService(userRepository, db)
	userService := services.NewUserService(userRepository, db)
	budgetService := services.NewBudgetService(budgetRepository, db)
	reportEntryService := services.NewReportService(recordEntryRepository, db)

	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	budgetController := controllers.NewBudgetController(budgetService)
	reportController := controllers.NewReportController(reportEntryService)

	app := fiber.New(fiber.Config{
		ErrorHandler: custom_errors.ErrorHandler,
	})
	app.Use(cors.New())

	authRouting(app, authController)
	userRouting(app, userController)
	budgetRouting(app, budgetController)
	reportRouting(app, reportController)

	return app, nil
}
