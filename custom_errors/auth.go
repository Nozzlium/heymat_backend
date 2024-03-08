package custom_errors

import (
	"database/sql"

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
				Code:    BAD_REQUEST_CODE,
				Status:  BAD_REQUEST_STATUS,
				Message: INVALID_CREDENTIALS,
			}
		}
	}
	return err
}
