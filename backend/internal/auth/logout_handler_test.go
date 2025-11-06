package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
)

func TestLogoutUserHandler(t *testing.T) {
	tokenString := "token"
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPut, "/auth/logout", nil)
	ctx := context.WithValue(req.Context(), middleware.CtxTokenString, tokenString)
	req = req.WithContext(ctx)
	if err != nil {
		t.Fatalf("unexpected error = %v", err)
	}

	repo := &mockTokenRepo{}
	h := NewHandler(&authService{
		tokenSvc: &token.TokenService{
			Repo: repo,
		},
	})

	h.LogoutUser(rr, req)

	if statusCode := rr.Result().StatusCode; statusCode != http.StatusOK {
		t.Errorf("expected status code = %d, got status code = %d", http.StatusOK, statusCode)
	}

	if !repo.DeleteByTokenStringCalled {
		t.Error("expected repo.DeleteByTokenString to be called")
	}
}
