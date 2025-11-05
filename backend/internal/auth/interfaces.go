package auth

import (
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
)

type authService struct {
	userSvc  *user.UserService
	tokenSvc *token.TokenService
}
