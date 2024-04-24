package auth

import (
	"errors"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nozzlium/heymat_backend/lib"
)

var (
	ErrUnauthorized = errors.New(
		"unauthorized",
	)
	ErrBadToken = errors.New(
		"bad token",
	)
	ErrUnknownClaims = errors.New(
		"unknown claims",
	)
)

type UserAccountResult struct {
	ID    uint64
	Email string
}

func AuthMiddleware(
	ctx *fiber.Ctx,
) error {
	tokenValue := ctx.Get(
		"Authorization",
		"",
	)
	if tokenValue == "" {
		return ctx.Next()
	}

	tokenElements := strings.Split(
		tokenValue,
		" ",
	)
	if len(tokenElements) < 2 {
		return lib.WriteResponse(
			ctx,
			fiber.StatusBadRequest,
			ErrBadToken.Error(),
		)
	}

	tokenPrefix := tokenElements[0]
	if tokenPrefix != "Bearer" {
		return lib.WriteResponse(
			ctx,
			fiber.StatusBadRequest,
			ErrBadToken.Error(),
		)
	}

	tokenString := tokenElements[1]
	token, err := lib.ParseJwtToken(
		tokenString,
	)
	if err != nil {
		log.Println(err)
		return lib.WriteResponse(
			ctx,
			fiber.StatusUnauthorized,
			ErrUnauthorized.Error(),
		)
	} else if claims, ok := token.Claims.(*lib.AuthClaims); ok {
		ctx.Context().SetUserValue(lib.USER_DATA, claims)
		return ctx.Next()
	} else {
		err = ErrUnknownClaims
		log.Println(err)
		return lib.WriteResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
}
