package token

import (
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

type TokenHandler struct{}

func NewHandler() *TokenHandler {
	return &TokenHandler{}
}

func RegisterRoutes(router *httprouter.Router) {
	h := NewHandler()
	router.Handler(http.MethodPut, "/tokens/refresh", middleware.RequireAuthentication(h.NewAccessToken, "REFRESH"))
}
