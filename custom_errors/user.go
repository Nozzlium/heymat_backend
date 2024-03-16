package custom_errors

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

const (
	EMAIL_ALREADY_EXISTS    = "email sudah terdaftar"
	USERNAME_ALREADY_EXISTS = "username sudah terdaftar"
)

var messages = map[string]string{
	"email":    EMAIL_ALREADY_EXISTS,
	"username": USERNAME_ALREADY_EXISTS,
}

func ParseRegisterError(err error) error {
	if driverErr, ok := err.(*pq.Error); ok {
		switch driverErr.Code.Name() {
		case "unique_violation":
			return getUniqueViolationError(driverErr)
		}
	}
	return err
}

func getUniqueViolationError(driverErr *pq.Error) error {
	constraint := strings.Split(driverErr.Constraint, "_")
	return &HttpError{
		Code:    fiber.ErrBadRequest.Code,
		Message: fiber.ErrBadRequest.Message,
		Detail:  messages[constraint[1]],
	}
}