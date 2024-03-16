package controllers

import "github.com/gofiber/fiber/v2"

type ReportController interface {
	Create(ctx *fiber.Ctx) error
}
