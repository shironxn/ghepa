package util

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field  string `json:"field"`
	Errors string `json:"error"`
}

func Validate(validate *validator.Validate, data interface{}) []ValidationError {
	err := validate.Struct(data)
	if err != nil {
		errors := make([]ValidationError, 0)

		for _, validationErr := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field:  validationErr.StructField(),
				Errors: validationErr.Tag(),
			})
		}

		return errors
	}

	return nil
}
