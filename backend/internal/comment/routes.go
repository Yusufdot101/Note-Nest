package comment

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/julienschmidt/httprouter"
)

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

	router.Handler(http.MethodPost, "/notes/:noteid/comments", middleware.RequireAccess(h.newComment))
}
