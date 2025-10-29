package user

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	repo := &mockRepo{}
	h := userHandler{
		svc: &UserService{
			repo: repo,
		},
	}

	rr := httptest.NewRecorder()
	msg := `{
		"name":"Yusuf",
		"email": "ym@gmail.com", 
		"password": "12345678"
	}`
	rr.Body.Write([]byte(msg))
	req, err := http.NewRequest(http.MethodPost, "/users/signup", rr.Body)
	if err != nil {
		t.Fatalf("unexpected error = %v", err)
	}
	h.RegisterUser(rr, req, nil)
	if !repo.insertUserCalled {
		t.Fatal("expected repo.insertUser to be called")
	}
}
