package project

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB, rdb *redis.Client) {
	h := NewHandler(&ProjectService{
		Repo: &Repository{
			DB: DB,
		},
		RDB: rdb,
	})

	router.Handler(http.MethodPost, "/projects", middleware.RequireAccess(h.NewProject))
	router.Handler(http.MethodGet, "/projects", middleware.Authenticate(h.GetProjects))

	router.Handler(http.MethodGet, "/projects/:id", middleware.Authenticate(h.GetProject))
	router.Handler(http.MethodPatch, "/projects/:id/visibility", middleware.RequireAccess(h.updateProjectVisibility))
	router.Handler(http.MethodPatch, "/projects/:id/color", middleware.RequireAccess(h.updateProjectColor))
	router.Handler(http.MethodDelete, "/projects/:id", middleware.RequireAccess(h.DeleteProject))

	router.Handler(http.MethodPatch, "/projects/:id", middleware.RequireAccess(h.UpdateProject))
}
