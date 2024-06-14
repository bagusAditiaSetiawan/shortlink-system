package exception

type ErrorBodyException struct {
	Message string
}

func (err ErrorBodyException) Error() string {
	return err.Message
}

func NewErrorBodyException(message string) *ErrorBodyException {
	return &ErrorBodyException{
		Message: message,
	}
}
