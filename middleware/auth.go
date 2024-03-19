package middleware

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nozzlium/heymat_backend/helper"
)

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
		return err
	} else if claims, ok := token.Claims.(*helper.AuthClaims); ok {
		ctx.Context().SetUserValue("userData", claims)
		return ctx.Next()
	} else {
		fmt.Println(claims)
		fmt.Println("2B", "9S")
		return errors.New("unknown claims")
	}
}
