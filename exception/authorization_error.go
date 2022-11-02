package exception

type AuthorizationError struct {
	StatusCode int
	message    string
}

func NewAuthorizationError(message string) *AuthorizationError {
	return &AuthorizationError{
		StatusCode: 403,
		message:    message,
	}
}

func (b AuthorizationError) Error() string {
	return b.message
}
