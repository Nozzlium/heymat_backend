package lib

import (
	"github.com/nozzlium/heymat_backend/controllers"
	"github.com/nozzlium/heymat_backend/repositories"
	"github.com/nozzlium/heymat_backend/services"
)

var DB = initDB()

var (
	UserRepository   = repositories.NewUserRepository()
	BudgetRepository = repositories.NewBudgetRepository()
)

var (
	AuthRepository = services.NewAuthService(UserRepository, DB)
	UserService    = services.NewUserService(UserRepository, DB)
	BudgetService  = services.NewBudgetService(BudgetRepository, DB)
)

var (
	AuthController   = controllers.NewAuthController(AuthRepository)
	UserController   = controllers.NewUserController(UserService)
	BudgetController = controllers.NewBudgetController(BudgetService)
)
