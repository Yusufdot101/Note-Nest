package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/golang-jwt/jwt/v4"
)

func TestNewAccessToken(t *testing.T) {
	userID := 1
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPut, "/auth/refreshtoken", nil)
	ctx := context.WithValue(req.Context(), middleware.CtxUserIDKey, userID)
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

	h.NewAccessToken(rr, req)

	if statusCode := rr.Result().StatusCode; statusCode != http.StatusOK {
		t.Errorf("expected status code = %d, got status code = %d", http.StatusOK, statusCode)
	}

	var input struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.Unmarshal(rr.Body.Bytes(), &input); err != nil {
		t.Fatalf("unexpected error = %v", err)
	}

	tokenString := input.AccessToken

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := token.ValidateJWT(tokenString, jwtSecret)
	if err != nil {
		t.Fatalf("unexpected error = %v", err)
	}

	claims := token.Claims.(jwt.MapClaims)
	issuer, ok := claims["iss"].(string)
	if !ok || issuer != os.Getenv("JWT_ISSUER") {
		// custom_errors.InvalidAuthenticationTokenErrorResponse(w)
		t.Error("invalid issuer")
	}

	subStr, ok := claims["sub"].(string)
	if !ok || subStr == "" {
		t.Error("invalid subject")
	}
	subInt, err := strconv.Atoi(subStr)
	if err != nil {
		t.Error("invalid subject format")
	}

	if subInt != userID {
		t.Errorf("expected subject = %d, got subject = %d", userID, subInt)
	}
}
