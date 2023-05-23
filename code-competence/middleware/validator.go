package middleware

import (
	"code-competence/models"
	"errors"

	"github.com/go-playground/validator/v10"
)

func EmailValidation(email string) error {
	v := validator.New()

	err := v.Var(email, "required,email")

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func PostUserValidation(user models.User) error {
	v := validator.New()

	err := v.Struct(user)

	if err != nil {
		return errors.New(err.Error())
	}

	validateEmail := EmailValidation(user.Email)

	if validateEmail != nil {
		return validateEmail
	}

	return nil
}
func PostItemValidation(items models.Items) error {
	v := validator.New()

	err := v.Struct(items)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
