package demo

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type RoleProfile struct {
	UserName string `validate:"gt=0"`
	Religion string `validate:"isErisClut"`
}

func (u *RoleProfile) validate() error {
	validate := validator.New()
	validate.RegisterValidation("isErisClut", validateFunc)
	err := validate.Struct(u)
	return err
}

func validateFunc(f1 validator.FieldLevel) bool {
	return !strings.Contains(f1.Field().String(), "Eris")
}
