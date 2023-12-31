package util

import (
	"event-planning-app/internal/response"
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func Validate(validate *validator.Validate, data interface{}) []response.ValidationError {
	err := validate.Struct(data)

	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		errors := make([]response.ValidationError, 0)

		for _, validationErr := range err.(validator.ValidationErrors) {
			translatedErr := fmt.Errorf(validationErr.Translate(trans))
			errors = append(errors, response.ValidationError{
				Field:  validationErr.StructField(),
				Errors: translatedErr.Error(),
			})
		}

		return errors
	}

	return nil
}
