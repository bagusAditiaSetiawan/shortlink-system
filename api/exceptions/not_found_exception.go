package exception

type NotFoundHandler struct {
	Message string
}

func (notFoundHandler NotFoundHandler) Error() string {
	return notFoundHandler.Message
}

func NewNotFoundHandler(err string) NotFoundHandler {
	return NotFoundHandler{Message: err}
}
