package user

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (us *UserService) registerUser(v *validator.Validator, name, email, password string) (string, string, error) {
	validatePassword(v, strings.TrimSpace(password)) // prevent passwords like , "        ", from being accepted
	validateName(v, strings.TrimSpace(name))
	validateEmail(v, strings.TrimSpace(email))
	if !v.IsValid() {
		return "", "", validator.ErrFailedValidation
	}
	u := &User{
		Name:  name,
		Email: email,
	}
	err := u.Password.Set(password)
	if err != nil {
		return "", "", err
	}

	err = us.repo.insertUser(u)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := token.CreateJWT("REFRESH", u.ID)
	if err != nil {
		return "", "", err
	}
	accessToken, err := token.CreateJWT("ACCESS", u.ID)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}
