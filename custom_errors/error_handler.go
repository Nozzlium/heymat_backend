package custom_errors

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if httpError, ok := err.(*HttpError); ok {
		return ctx.JSON(fiber.Map{
			"code":   httpError.Code,
			"status": httpError.Status,
			"error": fiber.Map{
				"message": httpError.Message,
			},
		}, "application/json")
	}
	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusInternalServerError,
		"status": "internal server error",
		"error": fiber.Map{
			"message": err.Error(),
		},
	}, "application/json")
}
