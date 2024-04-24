package lib

import "github.com/gofiber/fiber/v2"

type Response struct {
	Data interface{} `json:"data"`
}

func WriteResponse(
	ctx *fiber.Ctx,
	status int,
	data interface{},
) error {
	return ctx.Status(status).
		JSON(Response{
			Data: data,
		})
}
