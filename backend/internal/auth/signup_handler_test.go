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

func TestSignupHandler(t *testing.T) {
	tests := []struct {
		name                  string
		username              string
		email                 string
		password              string
		wantStatusCode        int
		wantInsertUserCalled  bool
		wantInsertTokenCalled bool
		wantErrors            bool
	}{
		{
			name:                  "valid inputs",
			username:              "yusuf",
			email:                 "ym@gmail.com",
			password:              "12345678",
			wantInsertUserCalled:  true,
			wantInsertTokenCalled: true,
			wantStatusCode:        http.StatusCreated,
		},
		{
			name:                 "missing name",
			username:             "",
			email:                "ym@gmail.com",
			password:             "12345678",
			wantInsertUserCalled: false,
			wantStatusCode:       http.StatusBadRequest,
			wantErrors:           true,
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
			"name": "%s",
			"email": "%s",
			"password": "%s"
			}`, test.username, test.email, test.password)

			req, err := http.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(msg))
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			h.SignupUser(rr, req)

			if status := rr.Result().StatusCode; status != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, status)
			}

			if userRepo.InsertUserCalled != test.wantInsertUserCalled {
				t.Fatalf("expected userRepo.InsertUserCalled = %v, got userRepo.InsertUserCalled = %v", test.wantInsertUserCalled, userRepo.GetUserByEmailCalled)
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
