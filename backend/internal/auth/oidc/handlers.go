package oidc

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/coreos/go-oidc/v3/oidc"
)

func (h *oidcHandler) googleBegin(w http.ResponseWriter, r *http.Request) {
	state := token.GenerateRandomTokenString()
	nonce := token.GenerateRandomTokenString()

	c := &http.Cookie{
		Name:     "state",
		Value:    state,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, c)

	c = &http.Cookie{
		Name:     "nonce",
		Value:    nonce,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, c)

	var err error
	GoogleConfig, GoogleProvider, err = googleConfig(GoogleClientID, GoogleClientSecret, GoogleRedirectURL)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	url := GoogleConfig.AuthCodeURL(state, oidc.Nonce(nonce))
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *oidcHandler) googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state, err := r.Cookie("state")
	if err != nil || r.URL.Query().Get("state") != state.Value {
		customerrors.BadRequestErrorResponse(w, errors.New("state mismatch"))
		return
	}

	nonce, err := r.Cookie("nonce")
	if err != nil || r.URL.Query().Get("state") != state.Value {
		customerrors.BadRequestErrorResponse(w, errors.New("missing nonce"))
		return
	}

	if GoogleConfig == nil || GoogleProvider == nil {
		customerrors.ServerErrorResponse(w, errors.New("google oidc not configured"))
		return
	}

	ctx := context.Background()

	user, err := googleCallback(ctx, r.URL.Query().Get("code"), GoogleConfig, GoogleProvider, nonce.Value)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	refreshToken, _, err := h.svc.findOrInsertUser(user)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	ttl, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRATION_TIME"))
	if err != nil {
		customerrors.ServerErrorResponse(w, errors.New("invalid refresh token expiration time"))
		return
	}
	err = utilities.SetTokenCookie(w, "REFRESH", refreshToken, "/auth", ttl)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	http.Redirect(w, r, "http://127.0.0.1:3000/notes", http.StatusFound)
}
