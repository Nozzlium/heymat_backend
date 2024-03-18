package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/requestbodies"
	"github.com/nozzlium/heymat_backend/services"
)

type BudgetControllerImpl struct {
	BudgetService services.BudgetService
}

func (controller *BudgetControllerImpl) Create(ctx *fiber.Ctx) error {
	body := requestbodies.Budget{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	controller.BudgetService.Create(ctx.Context(), entities.Budget{})
}
