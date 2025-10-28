package user

import (
	"regexp"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

var EmailRX = regexp.MustCompile(
	"^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
)

type User struct {
	ID          int
	Name, Email string
	Password    password
}

type password struct {
	plaintext *string // it is easier to check if the password was given using plaintext == nil
}

func (p *password) Set(plaintextPassword string) error {
	p.plaintext = &plaintextPassword
	return nil
}

func validateName(v *validator.Validator, name string) {
	v.CheckAddError(name != "", "name", "must be provided")
}

func validatePassword(v *validator.Validator, plaintextPassword string) {
	v.CheckAddError(plaintextPassword != "", "password", "must be provided")
	v.CheckAddError(len(plaintextPassword) >= 8, "password", "cannot be less than 8 characters long")
}

func validateEmail(v *validator.Validator, email string) {
	v.CheckAddError(email != "", "email", "must be provided")
	v.CheckAddError(validator.Matches(EmailRX, email), "email", "must be valid email address")
}
