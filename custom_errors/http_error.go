package custom_errors

const (
	BAD_REQUEST_CODE           = 400
	INTERNAL_SERVER_ERROR_CODE = 500
)

const (
	BAD_REQUEST_STATUS           = "bad request"
	INTERNAL_SERVER_ERROR_STATUS = "internal server error"
)

type HttpError struct {
	Code    int
	Status  string
	Message string
}

func (err *HttpError) Error() string {
	return err.Message
}
