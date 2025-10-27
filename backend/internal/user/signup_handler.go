package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type userHandler struct {
	svc *UserService
}

func NewHandler(s *UserService) *userHandler {
	return &userHandler{
		svc: s,
	}
}

func RegisterRoutes(router *httprouter.Router) {
	h := NewHandler(&UserService{})
	router.GET("/users/signup", h.RegisterUser)
}

func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	msg := "hello mate\n"
	_, err := w.Write([]byte(msg))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
