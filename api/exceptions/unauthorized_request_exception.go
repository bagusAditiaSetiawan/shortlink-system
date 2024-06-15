package exception

type UnauthorizedRequestException struct {
	Message string
}

func (UnauthorizedRequestException UnauthorizedRequestException) Error() string {
	return UnauthorizedRequestException.Message
}

func NewUnauthorizedRequestException(err string) UnauthorizedRequestException {
	return UnauthorizedRequestException{Message: err}
}
