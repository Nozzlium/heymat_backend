package custom_errors

type HttpError struct {
	Code    int
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (err *HttpError) Error() string {
	return err.Detail
}
