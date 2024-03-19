package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nozzlium/heymat_backend/entities"
)

var tempSignKey = "temp"

type AuthClaims struct {
	ID uint32 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(user entities.User) (string, error) {
	claims := AuthClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tempSignKey))
}
