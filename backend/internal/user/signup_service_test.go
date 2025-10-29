package user

import (
	"errors"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

type mockRepo struct {
	insertUserCalled bool
}

func (mr *mockRepo) insertUser(u *User) error {
	mr.insertUserCalled = true
	return nil
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name                              string
		userName, userEmail, userPassword string
		expectedErr                       error
	}{
		{
			name:         "valid input",
			userName:     "Yusuf",
			userEmail:    "ym@gmail.com",
			userPassword: "12345678",
		},
		{
			name:         "validation error",
			userName:     "",
			userEmail:    "ym@gmail.com",
			userPassword: "12345678",
			expectedErr:  validator.ErrFailedValidation,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &mockRepo{}
			svc := UserService{
				repo: repo,
			}
			v := validator.NewValidator()
			err := svc.registerUser(v, test.userName, test.userEmail, test.userPassword)
			if test.expectedErr != nil {
				if err == nil {
					t.Fatalf("expected error = %v, got none", test.expectedErr)
				}
				if !errors.Is(err, test.expectedErr) {
					t.Fatalf("expected error = %v, got error = %v", test.expectedErr, err)
				}
				return
			} else {
				if err != nil {
					t.Fatalf("expected no error, got error = %v", err)
				}
			}

			if !repo.insertUserCalled {
				t.Fatal("expected repo.insertUserCalled = true, got repo.insertUserCalled = false")
			}
		})
	}
}
