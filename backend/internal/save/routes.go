package save

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB, rdb *redis.Client) {
	h := newHandler(&saveService{
		repo: &repository{
			DB: DB,
		},
		RDB: rdb,
	})

	router.Handler(http.MethodPost, "/notes/:id/save", middleware.RequireAccess(h.saveNote))
	router.Handler(http.MethodDelete, "/notes/:id/save", middleware.RequireAccess(h.unsaveNote))
	router.Handler(http.MethodGet, "/notes/:id/save", middleware.RequireAccess(h.noteIsSaved))
	router.Handler(http.MethodGet, "/saved/notes", middleware.RequireAccess(h.savedNotes))
}
