package app

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func ConfigureRouter(router *httprouter.Router, DB *sql.DB) http.Handler {
	router.NotFound = http.HandlerFunc(custom_errors.NotFoundErrorResponse)
	router.MethodNotAllowed = http.HandlerFunc(custom_errors.MethodNotAllowedErrorResponse)

	user.RegisterRoutes(router, DB)
	token.RegisterRoutes(router)
	return middleware.EnableCORS(router)
}
