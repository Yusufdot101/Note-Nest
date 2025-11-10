package note

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := newHandler(&NoteService{
		Repo: &Repository{
			DB: DB,
		},
		ProjectSvc: &project.ProjectService{
			Repo: &project.Repository{
				DB: DB,
			},
		},
	})
	router.Handler(http.MethodPost, "/projects/:projectid/notes", middleware.RequireAccess(h.newNote))
	router.Handler(http.MethodGet, "/projects/:projectid/notes", middleware.RequireAccess(h.getNotes))
}
