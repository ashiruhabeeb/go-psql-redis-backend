package validator

import "github.com/go-playground/validator/v10"

var val = validator.New()

type ErrorResponse struct {
	Field	string	`json:"field"`
	Tag		string	`json:"tag"`
	Value	string	`json:"value"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse

	if err := val.Struct(payload); err != nil {
		for _, err := range err.(validator.ValidationErrors){
			var e ErrorResponse
			e.Field = err.StructNamespace()
			e.Tag = err.Tag()
			e.Value = err.Param()

			errors = append(errors, &e)
		}
	}
	return errors
}

func Validate(obj interface{}) error {
	return validator.New().Struct(obj)
}