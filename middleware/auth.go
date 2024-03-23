package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nozzlium/heymat_backend/custom_errors"
	"github.com/nozzlium/heymat_backend/entities"
	"github.com/nozzlium/heymat_backend/helper"
	"github.com/nozzlium/heymat_backend/lib"
)

func userMatch(entity entities.UserAccount, claim helper.AuthClaims) bool {
	return entity.ID == claim.ID && entity.Username == claim.Username && entity.Email == claim.Email
}

func AuthMiddleware(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization", "")
	token, err := jwt.ParseWithClaims(
		tokenString,
		&helper.AuthClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte("temp"), nil
		},
	)
	if err != nil {
		return custom_errors.GetUnauthorizedError()
	} else if claims, ok := token.Claims.(*helper.AuthClaims); ok {
		user, err := lib.UserRepository.FindById(ctx.Context(), lib.DB, claims.ID)
		if err != nil || !userMatch(user, *claims) {
			return custom_errors.GetUnauthorizedError()
		}
		ctx.Context().SetUserValue("userData", claims)
		return ctx.Next()
	} else {
		return errors.New("unknown claims")
	}
}
