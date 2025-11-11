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
	router.Handler(http.MethodGet, "/projects/:id", middleware.RequireAccess(h.GetProject))
	router.Handler(http.MethodDelete, "/projects/:id", middleware.RequireAccess(h.DeleteProject))
	router.Handler(http.MethodPatch, "/projects/:id", middleware.RequireAccess(h.UpdateProject))
}
