package validations

import (
	"api-users/models"
	"context"
	"errors"
	"regexp"
)

func ValidCreateUser(ctx context.Context, user models.Users) (string, error) {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if user.Name == "" {
		msg := "The name field is required!"
		newErr := errors.New(msg)
		return msg, newErr
	}
	if len(user.Name) < 2 || len(user.Name) > 40 {
		msg := "The name field must be between 2-40 chars!"
		newErr := errors.New(msg)
		return msg, newErr
	}
	if user.Email == "" {
		msg := "The email field is required!"
		newErr := errors.New(msg)
		return msg, newErr
	}
	if !regex.MatchString(user.Email) {
		msg := "The email field should be a valid email address!"
		newErr := errors.New(msg)
		return msg, newErr
	}
	if user.Password == "" {
		msg := "The password is required!"
		newErr := errors.New(msg)
		return msg, newErr
	}
	if len(user.Password) < 8 || len(user.Password) > 40 {
		msg := "The password field must be between 2-40 chars!"
		newErr := errors.New(msg)
		return msg, newErr
	}

	err := models.UniqueEmail(ctx, user)
	if err == nil {
		msg := "The email already exists"
		newErr := errors.New(msg)
		return msg, newErr
	}

	return "success", nil
}
