package custom_errors

import "github.com/gofiber/fiber/v2"

type HttpError struct {
	Code    int
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (err *HttpError) Error() string {
	return err.Detail
}

func GetUnauthorizedError() *HttpError {
	return &HttpError{
		Code:    fiber.ErrUnauthorized.Code,
		Message: fiber.ErrUnauthorized.Message,
		Detail:  "Anda belum masuk ke akun anda",
	}
}
