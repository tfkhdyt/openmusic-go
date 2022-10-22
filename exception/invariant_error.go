package exception

type BadRequestError struct {
	StatusCode int
	message    string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		StatusCode: 400,
		message:    message,
	}
}

func (b BadRequestError) Error() string {
	return b.message
}
