package note

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB, rdb *redis.Client) {
	repo, err := NewRepository(DB)
	if err != nil {
		panic(err)
	}

	h := newHandler(&NoteService{
		Repo: repo,
		ProjectSvc: &project.ProjectService{
			Repo: &project.Repository{
				DB: DB,
			},
			RDB: rdb,
		},
		RDB: rdb,
	})

	router.Handler(http.MethodPost, "/projects/:projectid/notes", middleware.RequireAccess(h.newNote))

	// specific note
	router.Handler(http.MethodGet, "/notes/:id", middleware.Authenticate(h.getNote))
	router.Handler(http.MethodDelete, "/notes/:id", middleware.RequireAccess(h.deleteNote))
	router.Handler(http.MethodPatch, "/notes/:id/content", middleware.RequireAccess(h.updateNoteTitleContent))
	router.Handler(http.MethodPatch, "/notes/:id/visibility", middleware.RequireAccess(h.updateNoteVisibility))
	router.Handler(http.MethodPatch, "/notes/:id/color", middleware.RequireAccess(h.updateNoteColor))

	// owner name
	router.Handler(http.MethodGet, "/notes/:id/username", middleware.Authenticate(h.getNoteOwner))

	// all projects and filtering/pagination
	router.Handler(http.MethodGet, "/notes", middleware.Authenticate(h.getNotes))
}
