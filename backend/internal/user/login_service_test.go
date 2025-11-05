package user

import (
	"errors"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func TestLoginUser(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		expectedErr error
	}{
		{
			name:     "valid credentials",
			email:    "ym@gmail.com",
			password: "12345678",
		},
		{
			name:        "invalid credentials",
			email:       "ym@gmail.com",
			password:    "aaaaaaaa",
			expectedErr: custom_errors.ErrInvalidCredentials,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &mockRepo{}
			svc := UserService{
				repo: repo,
			}
			v := validator.NewValidator()
			refreshToken, accessToken, err := svc.loginUser(v, test.email, test.password)
			if test.expectedErr != nil {
				if err == nil {
					t.Fatalf("expected error = %v, got none", test.expectedErr)
				}
				if !errors.Is(err, test.expectedErr) {
					t.Errorf("errors.Is: expected error = %v, got error = %v", test.expectedErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			if len(string(refreshToken)) == 0 || len(string(accessToken)) == 0 {
				t.Fatal("expected token to be returned")
			}
			if !repo.getUserByEmailCalled {
				t.Fatal("expected repo.getUserByEmail to be called")
			}
		})
	}
}
