/*
package auth provides authentication logic and handlers like logingin and signingup
*/
package auth

import (
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
)

type Mailer interface {
	Send(recipient, templateFile string, data any) error
}

type authService struct {
	userSvc  *user.UserService
	tokenSvc *token.TokenService
	mailer   Mailer
}

type authHandler struct {
	svc *authService
}

func newHandler(svc *authService) *authHandler {
	return &authHandler{
		svc: svc,
	}
}
