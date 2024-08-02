package utils

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

func ValidatorErrors(er error) map[string]string {
	fields := map[string]string{}

	for _, err := range er.(validator.ValidationErrors) {
		fields[err.Field()] = er.Error()
	}

	return fields
}
