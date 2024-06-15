package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func toResponse(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	return ctx.Status(statusCode).JSON(&fiber.Map{
		"errors": data,
	})
}

func ErrorHandlerException(ctx *fiber.Ctx, err error) error {
	errorMessages := []string{}
	if errNotFound := ErrorNotFoundHandler(err); errNotFound != nil {
		errorMessages = append(errorMessages, err.Error())
		return toResponse(ctx, http.StatusNotFound, errorMessages)
	}

	if errValidation := ErrorValidator(err); errValidation != nil {
		customValidator := CustomErrorValidator(errValidation)
		return toResponse(ctx, http.StatusBadRequest, customValidator)
	}

	if errBadRequestException := BadRequestExceptionHandler(err); errBadRequestException != nil {
		errorMessages = append(errorMessages, err.Error())
		return toResponse(ctx, http.StatusBadRequest, errorMessages)
	}
	if errUnauthorizedException := UnauthorizedExceptionHandler(err); errUnauthorizedException != nil {
		errorMessages = append(errorMessages, err.Error())
		return toResponse(ctx, http.StatusBadRequest, errorMessages)
	}

	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	if err != nil {
		errorMessages = append(errorMessages, err.Error())
		return toResponse(ctx, code, errorMessages)
	}

	return nil
}

func ErrorValidator(err error) validator.ValidationErrors {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		return exception
	} else {
		return nil
	}
}

func ErrorNotFoundHandler(err error) error {
	exception, ok := err.(NotFoundHandler)
	if ok {
		return exception
	} else {
		return nil
	}
}
func BadRequestExceptionHandler(err error) error {
	exception, ok := err.(BadRequestException)
	if ok {
		return exception
	} else {
		return nil
	}
}
func UnauthorizedExceptionHandler(err error) error {
	exception, ok := err.(UnauthorizedRequestException)
	if ok {
		return exception
	} else {
		return nil
	}
}
