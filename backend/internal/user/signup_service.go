package user

import (
	"github.com/Yusufdot101/note-nest/internal/validator"
)

type Repo interface {
	insertUser(u *User) error
}

type UserService struct {
	repo Repo
}

func (us *UserService) registerUser(v *validator.Validator, name, email, password string) error {
	validateName(v, name)
	validatePassword(v, password)
	validateEmail(v, email)
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
