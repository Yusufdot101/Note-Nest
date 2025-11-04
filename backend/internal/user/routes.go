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

	router.Handler(http.MethodPost, "/users/signup", middleware.EnableCORS(h.RegisterUser))
	router.Handler(http.MethodPost, "/users/login", middleware.EnableCORS(h.LoginUser))
	// for preflight, won't work otherwise
	router.Handler(http.MethodOptions, "/users/*any", middleware.EnableCORS(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
}
