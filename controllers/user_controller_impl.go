package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/requestbodies"
	"github.com/nozzlium/heymat_backend/services"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(
	userService services.UserService,
) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(ctx *fiber.Ctx) error {
	body := requestbodies.Register{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	resp, err := controller.UserService.Register(
		ctx.Context(),
		params.User{
			User: entities.User{
				Username: body.Username,
				Email:    body.Email,
				Password: body.Password,
			},
		},
	)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   resp,
	}, "application/json")
}
