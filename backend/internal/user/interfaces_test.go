package user

import (
	"errors"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func TestSetPassword(t *testing.T) {
	u := &User{
		ID:    0,
		Name:  "Yusuf",
		Email: "email@gmail.com",
	}
	tests := []struct {
		name            string
		actualPassword  string
		testingPassword string
		wantErr         bool
	}{
		{
			name:            "corret password",
			actualPassword:  "12345678",
			testingPassword: "12345678",
		},
		{
			name:            "incorrect password",
			actualPassword:  "12345678",
			testingPassword: "abcdefgh",
			wantErr:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := u.Password.Set(test.actualPassword)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}

			err = bcrypt.CompareHashAndPassword(u.Password.hash, []byte(test.testingPassword))
			if !test.wantErr && err == nil {
				return
			}
			if !test.wantErr && err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			if !errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				t.Fatalf("unexpected error = %v", err)
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name        string
		userName    string
		wantIsValid bool
	}{
		{
			name:        "valid name",
			userName:    "Yusuf",
			wantIsValid: true,
		},
		{
			name:        "empty name",
			wantIsValid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := validator.NewValidator()
			validateName(v, test.userName)
			if v.IsValid() != test.wantIsValid {
				t.Errorf("expected isValid = %v, got isValid = %v", test.wantIsValid, v.IsValid())
			}
		})
	}
}

func TestEmailName(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		wantIsValid bool
	}{
		{
			name:        "valid email",
			email:       "email@gmail.com",
			wantIsValid: true,
		},
		{
			name:        "empty email",
			wantIsValid: false,
		},
		{
			name:        "invalid email",
			email:       "email",
			wantIsValid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := validator.NewValidator()
			validateEmail(v, test.email)
			if v.IsValid() != test.wantIsValid {
				t.Errorf("expected isValid = %v, got isValid = %v", test.wantIsValid, v.IsValid())
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		wantIsValid bool
	}{
		{
			name:        "valid password",
			password:    "12345678",
			wantIsValid: true,
		},
		{
			name:        "empty password",
			wantIsValid: false,
		},
		{
			name:        "password shorter than 8 characters",
			password:    "123456",
			wantIsValid: false,
		},
		{
			name:        "password longer than 72 characters",
			password:    "12345678901234567890123456789012345678901234567890123456789012345678901234567890", // 80 chars
			wantIsValid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := validator.NewValidator()
			validatePassword(v, test.password)
			if v.IsValid() != test.wantIsValid {
				t.Errorf("expected isValid = %v, got isValid = %v", test.wantIsValid, v.IsValid())
			}
		})
	}
}
