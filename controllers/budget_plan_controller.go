package controllers

import "github.com/gofiber/fiber/v2"

type BudgetPlanController interface {
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
}
