package comment

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/julienschmidt/httprouter"
)

// RegisterRoutes registers the comments api routes
func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := newHandler(&commentService{
		repo: &repository{
			DB: DB,
		},
		noteSvc: &note.NoteService{
			Repo: &note.Repository{
				DB: DB,
			},
			ProjectSvc: &project.ProjectService{
				Repo: &project.Repository{
					DB: DB,
				},
			},
		},
	})

	router.Handler(http.MethodPost, "/notes/:id/comments", middleware.RequireAccess(h.newComment))
	router.Handler(http.MethodGet, "/notes/:id/comments", middleware.RequireAccess(h.getComments))
	router.Handler(http.MethodPatch, "/comments/:id", middleware.RequireAccess(h.updateComment))
	router.Handler(http.MethodDelete, "/comments/:id", middleware.RequireAccess(h.deleteComment))
}
