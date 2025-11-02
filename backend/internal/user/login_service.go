package user

import (
	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

type token string

func (us *UserService) loginUser(v *validator.Validator, email, password string) (token, error) {
	validateEmail(v, email)
	validatePassword(v, password)
	if !v.IsValid() {
		return "", validator.ErrFailedValidation
	}
	u, err := us.repo.getUserByEmail(email)
	if err != nil {
		return "", err
	}
	matches, err := u.Password.Matches(password)
	if err != nil {
		return "", err
	}
	if !matches {
		return "", custom_errors.ErrInvalidCredentials
	}
	// might do
	return token("token"), nil
}
