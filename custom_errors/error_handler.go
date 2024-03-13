package custom_errors

import (
	"github.com/gofiber/fiber/v2"
)

const ERROR = "error"

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if httpError, ok := err.(*HttpError); ok {
		ctx.Status(httpError.Code)
		return ctx.JSON(fiber.Map{
			"code":   httpError.Code,
			"status": ERROR,
			"error": fiber.Map{
				"message": httpError.Message,
				"detail":  httpError.Detail,
			},
		}, "application/json")
	}

	if err == fiber.ErrMethodNotAllowed {
		ctx.Status(fiber.ErrMethodNotAllowed.Code)
		return ctx.JSON(fiber.Map{
			"code":   fiber.ErrMethodNotAllowed.Code,
			"status": ERROR,
			"error": fiber.Map{
				"message": fiber.ErrMethodNotAllowed.Message,
				"detail":  err.Error(),
			},
		}, "application/json")
	}

	ctx.Status(fiber.ErrInternalServerError.Code)
	return ctx.JSON(fiber.Map{
		"code":   fiber.ErrInternalServerError.Code,
		"status": ERROR,
		"error": fiber.Map{
			"message": fiber.ErrInternalServerError.Message,
			"detail":  err.Error(),
		},
	}, "application/json")
}
