package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nozzlium/heymat_backend/entities"
)

var tempSignKey = "temp"

type AuthClaims struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(user entities.UserAccount) (string, error) {
	claims := AuthClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tempSignKey))
}
