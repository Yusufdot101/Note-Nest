package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	repo := &mockRepo{}
	h := userHandler{
		svc: &UserService{
			repo: repo,
		},
	}

	msg := `{
		"name":"Yusuf",
		"email": "ym@gmail.com", 
		"password": "12345678"
	}`
	req, err := http.NewRequest(http.MethodPost, "/users/signup", strings.NewReader(msg))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	if err != nil {
		t.Fatalf("unexpected error = %v", err)
	}

	h.RegisterUser(rr, req, nil)

	if !repo.insertUserCalled {
		t.Fatal("expected repo.insertUser to be called")
	}

	if repo.insertedUser.Name != "Yusuf" {
		t.Errorf("expected name = 'Yusuf', got = %q", repo.insertedUser.Name)
	}
	if repo.insertedUser.Email != "ym@gmail.com" {
		t.Errorf("expected email = 'ym@gmail.com', got = %q", repo.insertedUser.Email)
	}

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status code = %d, got status code = %d", http.StatusCreated, status)
	}
}
