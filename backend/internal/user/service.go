package user

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (us *UserService) NewUser(v *validator.Validator, name, email, password string) (*User, error) {
	trimmedPassword := strings.TrimSpace(password) // prevent passwords like , "        ", from being accepted
	trimmedName := strings.TrimSpace(name)
	trimmedEmail := strings.TrimSpace(email)

	validatePassword(v, trimmedPassword) // prevent passwords like , "        ", from being accepted
	validateName(v, trimmedName)
	validateEmail(v, trimmedEmail)
	if !v.IsValid() {
		return nil, validator.ErrFailedValidation
	}
	u := &User{
		Name:  trimmedEmail,
		Email: trimmedEmail,
	}
	err := u.Password.Set(trimmedPassword)
	if err != nil {
		return nil, err
	}

	err = us.Repo.InsertUser(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (us *UserService) VerifyAndGetUser(v *validator.Validator, email, password string) (*User, error) {
	trimmedPassword := strings.TrimSpace(password) // prevent passwords like , "        ", from being accepted
	trimmedEmail := strings.TrimSpace(email)
	validateEmail(v, trimmedEmail)
	validatePassword(v, trimmedPassword)
	if !v.IsValid() {
		return nil, validator.ErrFailedValidation
	}
	u, err := us.Repo.GetUserByEmail(trimmedEmail)
	if err != nil {
		return nil, err
	}
	matches, err := u.Password.Matches(trimmedPassword)
	if err != nil {
		return nil, err
	}
	if !matches {
		return nil, custom_errors.ErrInvalidCredentials
	}

	return u, nil
}
