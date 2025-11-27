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

	router.Handler(http.MethodGet, "/notes/:id/like", middleware.RequireAccess(h.noteIsLiked))
	router.Handler(http.MethodGet, "/comments/:id/like", middleware.RequireAccess(h.commentIsLiked))

	router.Handler(http.MethodPost, "/notes/:id/like", middleware.RequireAccess(h.addNoteLike))
	router.Handler(http.MethodPost, "/comments/:id/like", middleware.RequireAccess(h.addCommentLike))

	router.Handler(http.MethodDelete, "/notes/:id/like", middleware.RequireAccess(h.removeNoteLike))
	router.Handler(http.MethodDelete, "/comments/:id/like", middleware.RequireAccess(h.removeCommentLike))
}
