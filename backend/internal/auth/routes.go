package auth

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

type authHandler struct {
	svc *authService
}

func NewHandler(svc *authService) *authHandler {
	return &authHandler{
		svc: svc,
	}
}

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := NewHandler(&authService{
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

	router.Handler(http.MethodGet, "/auth/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := "this is test message\n"
		_, _ = w.Write([]byte(msg))
	}))
	router.Handler(http.MethodPost, "/auth/signup", http.HandlerFunc(h.SignupUser))
	router.Handler(http.MethodPut, "/auth/login", http.HandlerFunc(h.LoginUser))
	router.Handler(http.MethodPut, "/auth/refreshtoken", middleware.RequireRefresh(DB, h.NewAccessToken))
	router.Handler(http.MethodPut, "/auth/logout", middleware.RequireRefresh(DB, h.LogoutUser))
}
