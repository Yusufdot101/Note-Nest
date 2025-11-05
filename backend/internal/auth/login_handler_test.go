package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
)

func TestLoginHandler(t *testing.T) {
	tests := []struct {
		name                     string
		email                    string
		password                 string
		wantStatusCode           int
		wantGetUserByEmailCalled bool
		wantInsertTokenCalled    bool
		wantErrors               bool
	}{
		{
			name:                     "valid inputs",
			email:                    "ym@gmail.com",
			password:                 "12345678",
			wantGetUserByEmailCalled: true,
			wantInsertTokenCalled:    true,
			wantStatusCode:           http.StatusOK,
		},
		{
			name:                     "missing email",
			email:                    "",
			password:                 "12345678",
			wantGetUserByEmailCalled: false,
			wantStatusCode:           http.StatusBadRequest,
			wantErrors:               true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepo := &mockUserRepo{}
			tokenRepo := &mockTokenRepo{}
			h := NewHandler(&authService{
				userSvc: &user.UserService{
					Repo: userRepo,
				},
				tokenSvc: &token.TokenService{
					Repo: tokenRepo,
				},
			})

			msg := fmt.Sprintf(`{
			"email": "%s",
			"password": "%s"
			}`, test.email, test.password)

			req, err := http.NewRequest(http.MethodPut, "/auth/login", strings.NewReader(msg))
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			h.LoginUser(rr, req)

			if status := rr.Result().StatusCode; status != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, status)
			}

			if userRepo.GetUserByEmailCalled != test.wantGetUserByEmailCalled {
				t.Fatalf("expected userRepo.GetUserByEmailCalled = %v, got userRepo.GetUserByEmailCalled = %v", test.wantGetUserByEmailCalled, userRepo.GetUserByEmailCalled)
			}

			if tokenRepo.InsertTokenCalled != test.wantInsertTokenCalled {
				t.Fatalf("expected tokenRepo.InsertTokenCalled = %v, got tokenRepo.InsertTokenCalled = %v", test.wantInsertTokenCalled, tokenRepo.InsertTokenCalled)
			}

			var response struct {
				Token string `json:"token"`
				Error any    `json:"error"`
			}
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("unexpected error = %v", err)
			}

			if response.Token == "" && response.Error == nil {
				t.Errorf("expected token to be returned")
			}
			if test.wantErrors && response.Error == nil {
				t.Fatal("expected response.Errors got none")
			}
			if !test.wantErrors && response.Error != nil {
				t.Fatalf("expected response.Errors = none, got response.Error = %v", response.Error)
			}
		})
	}
}
