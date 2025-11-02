package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginUser(t *testing.T) {
	tests := []struct {
		name                     string
		email                    string
		password                 string
		expectedCode             int
		wantErrors               bool
		wantGetUserByEmailCalled bool
	}{
		{
			name:                     "valid credentials",
			email:                    "ym@gmail.com",
			password:                 "12345678",
			expectedCode:             http.StatusOK,
			wantGetUserByEmailCalled: true,
		},
		{
			name:                     "invalid credentials",
			email:                    "ym@gmail.com",
			password:                 "aaaaaaaa",
			expectedCode:             http.StatusBadRequest,
			wantErrors:               true,
			wantGetUserByEmailCalled: true,
		},
		{
			name:         "incorrect inputs",
			email:        "ym@gmail.com",
			password:     "", // empty password
			wantErrors:   true,
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &mockRepo{}
			svc := &UserService{
				repo: repo,
			}
			h := NewHandler(svc)

			msg := fmt.Sprintf(`{
			"email": "%s",
			"password": "%s"
			}`, test.email, test.password)

			req, err := http.NewRequest(http.MethodPost, "/user/login", strings.NewReader(msg))
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			// call the handler
			h.LoginUser(rr, req)

			if status := rr.Result().StatusCode; status != test.expectedCode {
				t.Errorf("expected status code = %d, got status code = %d", test.expectedCode, status)
			}

			if repo.getUserByEmailCalled != test.wantGetUserByEmailCalled {
				t.Fatalf("expected repo.getUserByEmail = %v, got repo.get = %v", test.wantGetUserByEmailCalled, repo.getUserByEmailCalled)
			}

			var response struct {
				Token string `json:"token"`
				Error any    `json:"error"`
			}
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("unexpected error = %v", err)
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
