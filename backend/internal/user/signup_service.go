package user

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

type Repo interface {
	insertUser(u *User) error
}

type UserService struct {
	repo Repo
}

func (us *UserService) registerUser(v *validator.Validator, name, email, password string) error {
	validatePassword(v, strings.TrimSpace(password)) // prevent passwords like , "        ", from being accepted
	validateName(v, strings.TrimSpace(name))
	validateEmail(v, strings.TrimSpace(email))
	if !v.IsValid() {
		return validator.ErrFailedValidation
	}
	u := &User{
		Name:  name,
		Email: email,
	}
	err := u.Password.Set(password)
	if err != nil {
		return err
	}

	return us.repo.insertUser(u)
}
