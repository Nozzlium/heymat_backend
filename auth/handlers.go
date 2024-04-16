package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/lib"
)

type RegisterRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type RegisterResponseBody struct {
	Code   uint8        `json:"code"`
	Status string       `json:"status"`
	Data   UserResponse `json:"data"`
}

type LoginResponseBody struct {
	Code   uint8         `json:"code"`
	Status string        `json:"status"`
	Data   LoginResponse `json:"data"`
}

func registerHandler(
	ctx *fiber.Ctx,
) error {
	body := RegisterRequestBody{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	resp, err := registerUser(
		ctx.Context(),
		UserAccount{
			Username: body.Username,
			Email:    body.Email,
			Password: body.Password,
		},
	)
	if err != nil {
		if err == ErrUserExists {
			return lib.WriteErrorResponse(
				fiber.StatusConflict,
				err,
				ctx,
			)
		}
		return err
	}

	return ctx.JSON(
		RegisterResponseBody{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   resp,
		},
		"application/json",
	)
}

func loginHandler(
	ctx *fiber.Ctx,
) error {
	body := LoginRequestBody{}
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	resp, err := login(
		ctx.Context(),
		UserAccount{
			Username: body.Identity,
			Email:    body.Identity,
			Password: body.Password,
		},
	)
	if err != nil {
		if err == ErrInvalidCredential {
			return lib.WriteErrorResponse(
				fiber.StatusUnauthorized,
				err,
				ctx,
			)
		}
		return err
	}

	return ctx.JSON(LoginResponseBody{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   resp,
	}, "application/json")
}
