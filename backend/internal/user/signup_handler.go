package user

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/jsonutil"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/julienschmidt/httprouter"
)

type userHandler struct {
	svc *UserService
}

func NewHandler(svc *UserService) *userHandler {
	return &userHandler{
		svc: svc,
	}
}

func RegisterRoutes(router *httprouter.Router, DB *sql.DB) {
	h := NewHandler(&UserService{
		repo: &Repository{DB: DB},
	})
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
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	err = h.svc.registerUser(v, input.Name, input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			custom_errors.FailedValidationErrorResponse(w, v.Errors)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = jsonutil.WriteJSON(w, jsonutil.Message{"message": "user created successfully"}, http.StatusCreated)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
