package custom_errors

import (
	"strings"

	"github.com/lib/pq"
)

const (
	EMAIL_ALREADY_EXISTS    = "email sudah terdaftar"
	USERNAME_ALREADY_EXISTS = "username sudah terdaftar"
)

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
	switch constraint[1] {
	case "email":
		return &HttpError{
			Code:    BAD_REQUEST_CODE,
			Status:  BAD_REQUEST_STATUS,
			Message: EMAIL_ALREADY_EXISTS,
		}
	case "username":
		return &HttpError{
			Code:    BAD_REQUEST_CODE,
			Status:  BAD_REQUEST_STATUS,
			Message: USERNAME_ALREADY_EXISTS,
		}
	}
	return driverErr
}
