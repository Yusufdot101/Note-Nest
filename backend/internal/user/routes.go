package user

import (
	"database/sql"

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
	// h := NewHandler(&UserService{
	// 	Repo: &Repository{DB: DB},
	// })
}
