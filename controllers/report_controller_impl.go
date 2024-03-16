package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/services"
)

type ReportControllerImpl struct {
	ReportService services.ReportService
}

func (controller *ReportControllerImpl) Create(ctx *fiber.Ctx) error {
}
