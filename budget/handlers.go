package budget

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/lib"
)

type BudgetPlanRequestBody struct {
	Title   string `json:"title"`
	Amount  uint64 `json:"amount"`
	Private bool   `json:"private"`
}

type BudgetPlanResponseBody struct {
	Code   uint8              `json:"code"`
	Status string             `json:"status"`
	Data   BudgetPlanResponse `json:"data"`
}

type BudgetPlanItemListResponseBody struct {
	Code   uint8                  `json:"code"`
	Status string                 `json:"status"`
	Data   BudgetPlanListResponse `json:"data"`
}

func createBudgetHandler(
	ctx *fiber.Ctx,
) error {
	claims, ok := ctx.Context().UserValue(lib.USER_DATA).(*lib.AuthClaims)
	if !ok {
		return lib.WriteResponse(
			ctx,
			fiber.StatusUnauthorized,
			"token tidak valid",
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

	return lib.WriteResponse(
		ctx,
		fiber.StatusOK,
		resp,
	)
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

	return lib.WriteResponse(
		ctx,
		fiber.StatusOK,
		resp,
	)
}

func getBudgetPlanItemsHandler(
	ctx *fiber.Ctx,
) error {
	claims, ok := ctx.Context().UserValue(lib.USER_DATA).(*lib.AuthClaims)
	if !ok {
		return lib.WriteResponse(
			ctx,
			fiber.StatusUnauthorized,
			"unauthorized",
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

	return lib.WriteResponse(
		ctx,
		fiber.StatusOK,
		resp,
	)
}
