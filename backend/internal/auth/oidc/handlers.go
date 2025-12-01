package oidc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/coreos/go-oidc/v3/oidc"
)

func (h *oidcHandler) googleBegin(w http.ResponseWriter, r *http.Request) {
	state, err := randString(16)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	nonce, err := randString(16)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

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

	GoogleConfig, GoogleProvider, err = googleConfig(GoogleClientID, GoogleClientSecret, GoogleRedirectURL)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	url := GoogleConfig.AuthCodeURL(state, oidc.Nonce(nonce))
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *oidcHandler) googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("state")
	if r.URL.Query().Get("state") != cookie.Value {
		customerrors.BadRequestErrorResponse(w, errors.New("state mismatch"))
		return
	}

	user, err := googleCallback(ctx, r.URL.Query().Get("code"), GoogleConfig, GoogleProvider)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	fmt.Fprintf(w, "Google user info:\nID: %s\nName: %s\nEmail: %s", user.ID, user.Name, user.Email)
}
