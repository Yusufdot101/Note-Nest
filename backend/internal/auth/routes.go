package auth

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/mailer"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

// RegisterRoutes registers the auth api routes
func RegisterRoutes(router *httprouter.Router, DB *sql.DB, mailer *mailer.Mailer) {
	h := newHandler(&authService{
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
		mailer: mailer,
	})

	router.Handler(http.MethodPost, "/auth/signup", http.HandlerFunc(h.signupUser))
	router.Handler(http.MethodPut, "/auth/login", http.HandlerFunc(h.loginUser))
	router.Handler(http.MethodPut, "/auth/refreshtoken", middleware.RequireRefresh(DB, h.newAccessToken))
	router.Handler(http.MethodPut, "/auth/logout", middleware.RequireRefresh(DB, h.logoutUser))
}
