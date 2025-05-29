package config

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitializeValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate
}
