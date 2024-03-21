package controllers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/services"
)

type ReportControllerImpl struct {
	ReportService services.ReportService
}

func NewReportController(
	reportService services.ReportService,
) *ReportControllerImpl {
	return &ReportControllerImpl{
		ReportService: reportService,
	}
}

func (controller *ReportControllerImpl) Create(ctx *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (controller *ReportControllerImpl) GetByYear(ctx *fiber.Ctx) error {
	claims, ok := ctx.Context().UserValue("userData").(*helper.AuthClaims)
	if !ok {
		return errors.New("unknown claims")
	}

	thisYear := time.Now().Year()
	year := ctx.QueryInt("year", thisYear)

	resp, err := controller.ReportService.GetByYear(ctx.Context(), params.ReportEntry{
		PageNo:   uint(ctx.QueryInt("pageNo", 1)),
		PageSize: uint(ctx.QueryInt("pageSize", 10)),
		RecordEntry: entities.ReportEntry{
			UserID: claims.ID,
			CreatedAt: time.Date(
				year,
				1,
				1,
				0,
				0,
				0,
				0,
				time.FixedZone("GMT+7", 7*60*60),
			),
		},
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
