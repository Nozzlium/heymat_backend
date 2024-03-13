package custom_errors

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const INVALID_CREDENTIALS = "username/email atau password salah"

var errs []error = []error{
	sql.ErrNoRows,
	bcrypt.ErrMismatchedHashAndPassword,
}

func ParseLoginError(err error) error {
	for _, correspErr := range errs {
		if err == correspErr {
			return &HttpError{
				Code:    fiber.ErrBadRequest.Code,
				Message: fiber.ErrBadRequest.Message,
				Detail:  INVALID_CREDENTIALS,
			}
		}
	}
	return err
}
