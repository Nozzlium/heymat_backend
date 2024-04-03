package lib

import "github.com/gofiber/fiber/v2"

func WriteResponse(
	code int,
	status string,
	data any,
	err string,
	ctx *fiber.Ctx,
) error {
	ctx.Response().SetStatusCode(code)
	return ctx.JSON(fiber.Map{
		"code":   code,
		"status": status,
		"data":   data,
		"error":  err,
	}, "application/json")
}

func WriteOKResponse(
	data any,
	ctx *fiber.Ctx,
) error {
	return WriteResponse(
		fiber.StatusOK,
		"success",
		data,
		"",
		ctx,
	)
}

func WriteErrorResponse(
	code int,
	err error,
	ctx *fiber.Ctx,
) error {
	return WriteResponse(
		code,
		"error",
		nil,
		err.Error(),
		ctx,
	)
}
