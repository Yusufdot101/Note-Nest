package oidc

import (
	"context"

	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
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

type userInfo struct {
	ProviderName string
	UserID       int
	Sub          string
	Name         string
	Email        string
}

type repo interface {
	insert(ui *userInfo) error
}

type oidcService struct {
	repo     repo
	userSvc  *user.UserService
	tokenSvc *token.TokenService
}

type oidcHandler struct {
	svc *oidcService
}

func newHandler(svc *oidcService) *oidcHandler {
	return &oidcHandler{
		svc: svc,
	}
}
