package user

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

type userHandler struct {
	svc *UserService
}

func NewHandler(svc *UserService) *userHandler {
	return &userHandler{
		svc: svc,
	}
}

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := NewHandler(&UserService{
		repo: &Repository{DB: DB},
	})

	router.Handler(http.MethodPost, "/users/signup", http.HandlerFunc(h.RegisterUser))
	router.Handler(http.MethodPost, "/users/login", http.HandlerFunc(h.LoginUser))
	router.Handler(http.MethodPut, "/users/logout", middleware.RequireAuthentication(h.LogoutUser))
}
