package oidc

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleRedirectURL = os.Getenv("GOOGLE_REDIRECT_URL")

	h := newHandler(&oidcService{
		repo: &repository{
			DB: DB,
		},
		userSvc: &user.UserService{
			Repo: &user.Repository{
				DB: DB,
			},
		},
		tokenSvc: &token.TokenService{
			Repo: &token.Repository{
				DB: DB,
			},
		},
	})

	router.HandlerFunc(http.MethodGet, "/auth/start/google", h.googleBegin)
	router.HandlerFunc(http.MethodGet, "/auth/callback/google", h.googleCallbackHandler)
}
