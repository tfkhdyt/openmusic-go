package exception

type InternalServerError struct {
	StatusCode int
	message    string
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{
		StatusCode: 500,
		message:    message,
	}
}

func (b InternalServerError) Error() string {
	return b.message
}
