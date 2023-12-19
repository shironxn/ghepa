package util

import (
	"event-planning-app/internal/response"

	"github.com/go-playground/validator/v10"
)

func Validate(validate *validator.Validate, data interface{}) []response.ValidationError {
	err := validate.Struct(data)
	if err != nil {
		errors := make([]response.ValidationError, 0)

		for _, validationErr := range err.(validator.ValidationErrors) {
			errors = append(errors, response.ValidationError{
				Field:  validationErr.StructField(),
				Errors: validationErr.Tag(),
			})
		}

		return errors
	}

	return nil
}
