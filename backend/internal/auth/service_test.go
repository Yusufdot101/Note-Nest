package auth

import (
	"testing"

	"github.com/Yusufdot101/note-nest/internal/token"
)

func TestGetTokens(t *testing.T) {
	svc := &authService{
		tokenSvc: &token.TokenService{
			Repo: &mockTokenRepo{},
		},
	}
	userID := 1
	refreshToken, accesstoken, err := svc.getTokens(userID)
	if err != nil {
		t.Fatalf("unexpected error = %v", err)
	}
	if refreshToken == "" || accesstoken == "" {
		t.Errorf("expected tokens to be returned")
	}
}
