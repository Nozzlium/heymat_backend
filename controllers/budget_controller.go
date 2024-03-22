package controllers

import "github.com/gofiber/fiber/v2"

type BudgetController interface {
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
}
