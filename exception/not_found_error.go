package exception

type NotFoundError struct {
	StatusCode int
	message    string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		StatusCode: 404,
		message:    message,
	}
}

func (b NotFoundError) Error() string {
	return b.message
}
