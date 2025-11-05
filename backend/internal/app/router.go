package app

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/auth"
	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func ConfigureRouter(router *httprouter.Router, DB *sql.DB) http.Handler {
	router.NotFound = http.HandlerFunc(custom_errors.NotFoundErrorResponse)
	router.MethodNotAllowed = http.HandlerFunc(custom_errors.MethodNotAllowedErrorResponse)

	auth.RegisterRoutes(router, DB)
	user.RegisterRoutes(router, DB)
	return middleware.EnableCORS(router)
}
