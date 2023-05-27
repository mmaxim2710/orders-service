package validations

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func InitValidator() {
	validate = validator.New()
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func UniversalValidation(body interface{}) (bool, []*ValidationError) {
	var errors []*ValidationError
	err := validate.Struct(body)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errors = append(errors, &ValidationError{
				Field:   e.Field(),
				Message: e.Field(),
			})
		}

		return false, errors
	}

	return true, nil
}
