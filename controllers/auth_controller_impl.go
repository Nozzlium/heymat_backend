package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/params"
	"github.com/nozzlium/heymat_backend/requestbodies"
	"github.com/nozzlium/heymat_backend/services"
)

type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthController(
	authService services.AuthService,
) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	body := requestbodies.Login{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	resp, err := controller.AuthService.Login(ctx.Context(), params.Auth{
		User: entities.User{
			Username: body.Identity,
			Email:    body.Identity,
			Password: body.Password,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   resp,
	}, "application/json")
}
