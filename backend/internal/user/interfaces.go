package user

import (
	"regexp"
	"time"

	"github.com/Yusufdot101/note-nest/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var EmailRX = regexp.MustCompile(
	"^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
)

const hashingCost = 12

type User struct {
	ID                   int
	CreatedAt, UpdatedAt time.Time
	Name, Email          string
	Password             password
}

type password struct {
	hash []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), hashingCost)
	if err != nil {
		return err
	}
	p.hash = hash
	return nil
}

func validateName(v *validator.Validator, name string) {
	v.CheckAddError(name != "", "name", "must be provided")
}

func validatePassword(v *validator.Validator, plaintextPassword string) {
	v.CheckAddError(plaintextPassword != "", "password", "must be provided")
	v.CheckAddError(len(plaintextPassword) >= 8, "password", "cannot be less than 8 characters long")
	v.CheckAddError(len(plaintextPassword) <= 72, "password", "cannot be more than 72 characters long") // bcrypt reasons
}

func validateEmail(v *validator.Validator, email string) {
	v.CheckAddError(email != "", "email", "must be provided")
	v.CheckAddError(validator.Matches(EmailRX, email), "email", "must be valid email address")
}
