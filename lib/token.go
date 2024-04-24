package lib

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	USER_DATA = "userData"
)

type AuthClaims struct {
	UserID   uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var SignKey string

func GenerateJwtToken(
	claims AuthClaims,
) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString(
		[]byte(SignKey),
	)
}

func ParseJwtToken(
	jwtToken string,
) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		jwtToken,
		&AuthClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(
				SignKey,
			), nil
		},
	)
}

func InitAuthClaims(
	userId uint64,
	username string,
	email string,
) AuthClaims {
	return AuthClaims{
		UserID:   userId,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().
					Add(time.Hour * 24),
			),
		},
	}
}

func SetSignKey() error {
	signKeyEnv, ok := os.LookupEnv(
		"SIGN_KEY",
	)
	if !ok {
		return errors.New(
			"sign key not found",
		)
	}
	SignKey = signKeyEnv
	return nil
}
