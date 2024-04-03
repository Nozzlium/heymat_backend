package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nozzlium/heymat_backend/lib"
)

var tempSignKey = "temp"

func generateJwtToken(
	user UserAccount,
) (string, error) {
	claims := lib.AuthClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().
					Add(time.Hour * 24),
			),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString(
		[]byte(tempSignKey),
	)
}
