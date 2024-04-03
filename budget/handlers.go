package budget

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/lib"
)

type BudgetPlanRequestBody struct {
	Title   string `json:"title"`
	Amount  uint64 `json:"amount"`
	Private bool   `json:"private"`
}

func createBudgetHandler(
	ctx *fiber.Ctx,
) error {
	claims, ok := ctx.Context().UserValue(lib.USER_DATA).(*lib.AuthClaims)
	if !ok {
		return errors.New(
			"unauthorized",
		)
	}

	body := BudgetPlanRequestBody{}
	ctx.BodyParser(&body)
	resp, err := createBudgetPlan(
		ctx.Context(),
		BudgetPlan{
			Title:   body.Title,
			Amount:  body.Amount,
			Private: body.Private,
			UserID:  claims.UserID,
		},
	)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "OK",
		"code":   fiber.StatusOK,
		"data":   resp,
	}, "application/json")
}

func getBudgetPlanByIdHandler(
	ctx *fiber.Ctx,
) error {
	var userId uint64 = 0
	claims, ok := ctx.Context().UserValue(lib.USER_DATA).(*lib.AuthClaims)
	if ok {
		userId = claims.UserID
	}

	budgetPlanId, err := ctx.ParamsInt(
		"budgetPlanId",
		0,
	)
	if err != nil {
		return err
	}

	resp, err := getBudgetPlanById(
		ctx.Context(),
		BudgetPlan{
			ID: uint64(
				budgetPlanId,
			),
			UserID: userId,
		},
	)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "OK",
		"code":   fiber.StatusOK,
		"data":   resp,
	}, "application/json")
}

func getBudgetPlanItemsHandler(
	ctx *fiber.Ctx,
) error {
	claims, ok := ctx.Context().UserValue(lib.USER_DATA).(*lib.AuthClaims)
	if !ok {
		return lib.WriteErrorResponse(
			fiber.StatusUnauthorized,
			errors.New("unauthorized"),
			ctx,
		)
	}

	param := BudgetPlanParam{
		PageNo: uint32(
			ctx.QueryInt("pageNo", 1),
		),
		PageSize: uint32(
			ctx.QueryInt(
				"pageSize",
				10,
			),
		),
		BudgetPlan: BudgetPlan{
			Title: ctx.Query(
				"keyword",
				"",
			),
			UserID: claims.UserID,
		},
	}
	resp, err := getBudgetPlanItemList(
		ctx.Context(),
		param,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "OK",
		"code":   fiber.StatusOK,
		"data":   resp,
	}, "application/json")
}
