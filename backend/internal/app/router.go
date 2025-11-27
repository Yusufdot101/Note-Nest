package app

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/auth"
	"github.com/Yusufdot101/note-nest/internal/comment"
	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/like"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/Yusufdot101/note-nest/internal/save"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func configureRouter(router *httprouter.Router, DB *sql.DB, rateLimiterEnabled bool, rateLimiterBurst int, rateLimiterRate float64) http.Handler {
	router.NotFound = http.HandlerFunc(customerrors.NotFoundErrorResponse)
	router.MethodNotAllowed = http.HandlerFunc(customerrors.MethodNotAllowedErrorResponse)

	auth.RegisterRoutes(router, DB)
	user.RegisterRoutes(router, DB)
	project.RegisterRoutes(router, DB)
	note.RegisterRoutes(router, DB)
	like.RegisterRoutes(router, DB)
	comment.RegisterRoutes(router, DB)
	save.RegisterRoutes(router, DB)

	return middleware.EnableCORS(middleware.RecoverPanic(middleware.RateLimit(router, rateLimiterEnabled, rateLimiterBurst, rateLimiterRate)))
}
