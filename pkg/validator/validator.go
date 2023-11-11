package validator

import "github.com/go-playground/validator/v10"

func Validate(obj interface{}) error {
	return validator.New().Struct(obj)
}