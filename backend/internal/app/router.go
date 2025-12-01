package app

import (
	"database/sql"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/auth"
	"github.com/Yusufdot101/note-nest/internal/auth/oidc"
	"github.com/Yusufdot101/note-nest/internal/comment"
	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/like"
	"github.com/Yusufdot101/note-nest/internal/mailer"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/Yusufdot101/note-nest/internal/save"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func configureRouter(router *httprouter.Router, cfg *config, DB *sql.DB) http.Handler {
	router.NotFound = http.HandlerFunc(customerrors.NotFoundErrorResponse)
	router.MethodNotAllowed = http.HandlerFunc(customerrors.MethodNotAllowedErrorResponse)

	m := mailer.New(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.Sender, cfg.SMTP.Username, cfg.SMTP.Password)
	auth.RegisterRoutes(router, DB, m)
	user.RegisterRoutes(router, DB)
	project.RegisterRoutes(router, DB)
	note.RegisterRoutes(router, DB)
	like.RegisterRoutes(router, DB)
	comment.RegisterRoutes(router, DB)
	save.RegisterRoutes(router, DB)
	oidc.RegisterRoutes(router)

	return middleware.EnableCORS(middleware.RecoverPanic(middleware.RateLimit(router, cfg.Limiter.Enabled, cfg.Limiter.Burst, cfg.Limiter.Rate)))
}
