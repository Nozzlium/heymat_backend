package lib

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	USER_DATA = "userData"
)

var (
	ErrUnauthorized = errors.New(
		"unauthorized",
	)
	ErrBadToken = errors.New(
		"bad token",
	)
	ErrUnknownClaims = errors.New("unknown claims")
)

type UserAccountResult struct {
	ID    uint64
	Email string
}

type AuthClaims struct {
	UserID   uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
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
		return WriteErrorResponse(fiber.StatusBadRequest, ErrBadToken, ctx)
	}

	tokenPrefix := tokenElements[0]
	if tokenPrefix != "Bearer" {
		return WriteErrorResponse(fiber.StatusBadRequest, ErrBadToken, ctx)
	}

	tokenString := tokenElements[1]
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AuthClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte("temp"), nil
		},
	)
	if err != nil {
		log.Println(err)
		return WriteErrorResponse(fiber.StatusUnauthorized, ErrUnauthorized, ctx)
	} else if claims, ok := token.Claims.(*AuthClaims); ok {
		userAccountResult, err := getUserByIdAndEmail(ctx.Context(), claims)
		if err != nil {
			log.Println(err)
			return WriteErrorResponse(fiber.StatusUnauthorized, ErrUnauthorized, ctx)
		}

		if !claimsDataMatchesUserData(claims, userAccountResult) {
			return WriteErrorResponse(fiber.StatusUnauthorized, ErrUnauthorized, ctx)
		}

		ctx.Context().SetUserValue(USER_DATA, claims)
		return ctx.Next()
	} else {
		err = ErrUnknownClaims
		log.Println(err)
		return WriteErrorResponse(fiber.StatusInternalServerError, err, ctx)
	}
}

func getUserByIdAndEmail(ctx context.Context, claims *AuthClaims) (UserAccountResult, error) {
	query := `
    select
      id,
      email
    from user_account
    where
      id = $1 and
      email = $2
  `
	var userAccountResult = UserAccountResult{}
	err := DB.QueryRowContext(ctx, query, claims.UserID, claims.Email).Scan(&userAccountResult.ID, &userAccountResult.Email)
	if err != nil {
		log.Println(err)
		return userAccountResult, err
	}
	return userAccountResult, nil
}

func claimsDataMatchesUserData(claims *AuthClaims, userAccountResult UserAccountResult) bool {
	return claims.UserID == userAccountResult.ID &&
		claims.Email == userAccountResult.Email
}
