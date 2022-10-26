package exception

type AuthenticationError struct {
	StatusCode int
	message    string
}

func NewAuthenticationError(message string) *AuthenticationError {
	return &AuthenticationError{
		StatusCode: 401,
		message:    message,
	}
}

func (b AuthenticationError) Error() string {
	return b.message
}
