package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nozzlium/heymat_backend/data/entities"
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
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().AddDate(0, 0, 1),
			},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tempSignKey)
}
