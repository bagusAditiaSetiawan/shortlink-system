package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func toResponse(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	return ctx.Status(statusCode).JSON(&fiber.Map{
		"data": data,
	})
}

func ErrorHandlerException(ctx *fiber.Ctx, err error) error {
	if errNotFound := ErrorNotFoundHandler(err); errNotFound != nil {
		return toResponse(ctx, http.StatusNotFound, err.Error())
	}

	if errValidation := ErrorValidator(err); errValidation != nil {
		customValidator := CustomErrorValidator(errValidation)
		return toResponse(ctx, http.StatusBadRequest, customValidator)
	}

	if errBadRequestException := BadRequestExceptionHandler(err); errBadRequestException != nil {
		return toResponse(ctx, http.StatusBadRequest, err.Error())
	}

	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	if err != nil {
		return toResponse(ctx, code, err.Error())
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
