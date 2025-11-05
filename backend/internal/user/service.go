package user

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (us *UserService) NewUser(v *validator.Validator, name, email, password string) (*User, error) {
	validatePassword(v, strings.TrimSpace(password)) // prevent passwords like , "        ", from being accepted
	validateName(v, strings.TrimSpace(name))
	validateEmail(v, strings.TrimSpace(email))
	if !v.IsValid() {
		return nil, validator.ErrFailedValidation
	}
	u := &User{
		Name:  name,
		Email: email,
	}
	err := u.Password.Set(password)
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
	validateEmail(v, email)
	validatePassword(v, password)
	if !v.IsValid() {
		return nil, validator.ErrFailedValidation
	}
	u, err := us.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	matches, err := u.Password.Matches(password)
	if err != nil {
		return nil, err
	}
	if !matches {
		return nil, custom_errors.ErrInvalidCredentials
	}

	return u, nil
}
