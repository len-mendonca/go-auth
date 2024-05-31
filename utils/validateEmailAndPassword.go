package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/len-mendonca/go-auth/models"
)

func ValidateEmailAndPassword(validate *validator.Validate, user *models.User) map[string]string {

	errors := make(map[string]string)

	if err := validate.Var(user.Email, "required,email"); err != nil {
		errors["Email"] = "invalid email format"
	}
	if err := validate.Var(user.Password, "required"); err != nil {
		errors["Password"] = "password required"
	}

	if len(errors) > 0 {
		return errors
	}
	return nil

}
