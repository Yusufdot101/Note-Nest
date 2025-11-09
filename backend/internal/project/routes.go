package project

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := NewHandler(&ProjectService{
		Repo: &Repository{
			DB: DB,
		},
	})

	router.Handler(http.MethodPost, "/projects", middleware.RequireAccess(h.NewProject))
	router.Handler(http.MethodGet, "/projects", middleware.RequireAccess(h.GetProjects))
}
