package validator

import "github.com/go-playground/validator/v10"

type UserValidator struct {
	validator *validator.Validate
}

func (cv *UserValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var Validator = &UserValidator{validator: validator.New()}
