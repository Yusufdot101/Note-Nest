package like

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := newHandler(&likeService{
		repo: &repo{DB: DB},
	})

	router.Handler(http.MethodPost, "/notes/:id/like", middleware.RequireAccess(h.addLike))
	router.Handler(http.MethodDelete, "/notes/:id/like", middleware.RequireAccess(h.removeLike))
}
