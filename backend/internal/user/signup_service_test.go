package user

import (
	"errors"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

type mockRepo struct {
	insertUserCalled     bool
	insertedUser         *User
	getUserByEmailCalled bool
	gotUser              *User
}

func (mr *mockRepo) insertUser(u *User) error {
	mr.insertUserCalled = true
	mr.insertedUser = u
	return nil
}

func (mr *mockRepo) getUserByEmail(email string) (*User, error) {
	mr.getUserByEmailCalled = true
	mr.gotUser = &User{
		ID:    0,
		Name:  "yusuf",
		Email: "ym@gmail.com",
	}
	err := mr.gotUser.Password.Set("12345678")
	if err != nil {
		return nil, err
	}
	return mr.gotUser, nil
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
			token, err := svc.registerUser(v, test.userName, test.userEmail, test.userPassword)
			if test.expectedErr != nil {
				if err == nil {
					t.Fatalf("expected error = %v, got none", test.expectedErr)
				}
				if !errors.Is(err, test.expectedErr) {
					t.Fatalf("expected error = %v, got error = %v", test.expectedErr, err)
				}
				if repo.insertUserCalled {
					t.Fatal("expected repo.insertUser not to be called on error")
				}
				return
			} else {
				if err != nil {
					t.Fatalf("expected no error, got error = %v", err)
				}
			}

			if len(string(token)) == 0 {
				t.Fatal("expected token to be returned")
			}
			if !repo.insertUserCalled {
				t.Fatal("expected repo.insertUserCalled = true, got repo.insertUserCalled = false")
			}
		})
	}
}
