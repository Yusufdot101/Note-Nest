package oidc

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router) {
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleRedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")

	GithubClientID = os.Getenv("GITHUB_CLIENT_ID")
	GithubClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	GithubRedirectURL = os.Getenv("GITHUB_CLIENT_ID")

	h := newHandler(&oidcService{})

	router.HandlerFunc(http.MethodGet, "/auth/start/google", h.googleBegin)
	router.HandlerFunc(http.MethodGet, "/auth/callback/google", h.googleCallbackHandler)
}
