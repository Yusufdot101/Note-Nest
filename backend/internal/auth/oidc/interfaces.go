package oidc

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

var GithubClientID, GithubClientSecret, GithubRedirectURL string

var (
	GoogleClientID, GoogleClientSecret, GoogleRedirectURL string
	GoogleConfig                                          *oauth2.Config
	GoogleProvider                                        *oidc.Provider
)

type UserInfo struct {
	ID    string
	Name  string
	Email string
}

type repo interface{}

type oidcService struct {
	repo repo
}

type oidcHandler struct {
	svc *oidcService
}

func newHandler(svc *oidcService) *oidcHandler {
	return &oidcHandler{
		svc: svc,
	}
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
