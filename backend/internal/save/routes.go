package save

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := newHandler(&saveService{
		repo: &repository{
			DB: DB,
		},
	})

	router.Handler(http.MethodPost, "/notes/:id/save", middleware.RequireAccess(h.saveNote))
	router.Handler(http.MethodGet, "/notes/:id/save", middleware.RequireAccess(h.noteIsSaved))
}
