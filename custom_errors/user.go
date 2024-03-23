package custom_errors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

const (
	EMAIL_OR_USERNAME_EXISTS = "email atau username sudah terdaftar"
)

func ParseRegisterError(err error) error {
	if driverErr, ok := err.(*pq.Error); ok {
		switch driverErr.Code.Name() {
		case "unique_violation":
			return &HttpError{
				Code:    fiber.ErrBadRequest.Code,
				Message: fiber.ErrBadGateway.Message,
				Detail:  EMAIL_OR_USERNAME_EXISTS,
			}
		}
	}
	return err
}
