package controllers

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Register(ctx *fiber.Ctx) error
}
