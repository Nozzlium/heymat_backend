package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/requestbodies"
	"github.com/nozzlium/heymat_backend/services"
)

type BudgetPlanControllerImpl struct {
	BudgetPlanService services.BudgetPlanService
}

func NewBudgetController(
	budgetPlanService services.BudgetPlanService,
) *BudgetPlanControllerImpl {
	return &BudgetPlanControllerImpl{
		BudgetPlanService: budgetPlanService,
	}
}

func (controller *BudgetPlanControllerImpl) Create(ctx *fiber.Ctx) error {
	claims, ok := ctx.Context().UserValue("userData").(*helper.AuthClaims)
	if !ok {
		return errors.New("unauthorized")
	}

	body := requestbodies.BudgetPlan{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	resp, err := controller.BudgetPlanService.Create(ctx.Context(), entities.BudgetPlan{
		UserID:  uint64(claims.ID),
		Title:   body.Title,
		Amount:  body.Amount,
		Private: body.Private,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "OK",
		"code":   fiber.StatusOK,
		"data":   resp,
	}, "application/json")
}

func (controller *BudgetPlanControllerImpl) Get(ctx *fiber.Ctx) error {
	claims, ok := ctx.Context().UserValue("userData").(*helper.AuthClaims)
	if !ok {
		return errors.New("unauthorized")
	}

	param := params.BudgetPlan{
		PageNo:   uint32(ctx.QueryInt("pageNo", 1)),
		PageSize: uint32(ctx.QueryInt("pageSize", 10)),
		BudgetPlan: entities.BudgetPlan{
			UserID: uint64(claims.ID),
		},
	}
	resp, err := controller.BudgetPlanService.Get(ctx.Context(), param)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"code":   fiber.StatusOK,
		"data":   resp,
	}, "application/json")
}
