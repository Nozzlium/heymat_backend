package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/requestbodies"
	"github.com/nozzlium/heymat_backend/services"
)

type BudgetControllerImpl struct {
	BudgetService services.BudgetService
}

func NewBudgetController(budgetService services.BudgetService) *BudgetControllerImpl {
	return &BudgetControllerImpl{
		BudgetService: budgetService,
	}
}

func (controller *BudgetControllerImpl) Create(ctx *fiber.Ctx) error {
	userData, ok := ctx.Context().UserValue("userData").(*helper.AuthClaims)
	if !ok {
		return errors.New("unknwon claims")
	}

	body := requestbodies.Budget{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	resp, err := controller.BudgetService.Create(ctx.Context(), entities.Budget{
		UserID:    userData.ID,
		Amount:    body.Amount,
		TimeCode:  helper.TruncateToMonth(body.Date),
		CreatedAt: body.Date,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   resp,
	}, "application/json")
}
