package user

import (
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/jsonutil"
	"github.com/Yusufdot101/note-nest/internal/validator"
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
	router.POST("/users/signup", h.RegisterUser)
}

func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := jsonutil.ReadJSON(w, r, &input)
	if err != nil {
		err = jsonutil.WriteJSON(w, jsonutil.Message{"error": err.Error()}, http.StatusBadRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}
	u := &User{
		Name:  input.Name,
		Email: input.Email,
	}
	err = u.Password.Set(input.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	v := validator.NewValidator()
	validateName(v, u.Name)
	validatePassword(v, *u.Password.plaintext)
	validateEmail(v, u.Email)
	if !v.IsValid() {
		err = jsonutil.WriteJSON(w, jsonutil.Message{"errors": v.Errors}, 400)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	err = jsonutil.WriteJSON(w, jsonutil.Message{"input": input}, 200)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
