package exception

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"shortlink-system/pkg/languages"
	"strings"
)

var ROLES = fiber.Map{
	"required": languages.REQUIRED,
	"gte":      languages.MIN,
	"email":    languages.EMAIL_INVALID,
}

func CustomErrorValidator(errs validator.ValidationErrors) []string {
	validationErrors := []string{}
	if errs != nil {
		for _, err := range errs {
			message := fmt.Sprintf("%s_%s", strings.ToUpper(err.Field()), ROLES[err.ActualTag()])
			validationErrors = append(validationErrors, message)
		}
	}

	return validationErrors
}
