package utils

import (
	validator "github.com/go-playground/validator/v10"
)

func ValidateStruct(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
