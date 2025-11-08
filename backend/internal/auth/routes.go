package auth

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

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

	router.Handler(http.MethodPost, "/auth/signup", http.HandlerFunc(h.SignupUser))
	router.Handler(http.MethodPut, "/auth/login", http.HandlerFunc(h.LoginUser))
	router.Handler(http.MethodPut, "/auth/refreshtoken", middleware.RequireRefresh(DB, h.NewAccessToken))
	router.Handler(http.MethodPut, "/auth/logout", middleware.RequireRefresh(DB, h.LogoutUser))
}
