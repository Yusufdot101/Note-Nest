package auth

import (
	"log"

	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (as *authService) getTokens(userID int) (string, string, error) {
	refreshToken, err := as.tokenSvc.NewToken(token.RANDOMSTRING, token.REFRESH, userID)
	if err != nil {
		return "", "", err
	}
	accessToken, err := as.tokenSvc.NewToken(token.JWT, token.ACCESS, userID)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (as *authService) registerUser(v *validator.Validator, name, email, password string) (string, string, error) {
	u, err := as.userSvc.NewUser(v, name, email, password)
	if err != nil {
		return "", "", err
	}

	as.sendEmail(u.Email, "user_welcome.tmpl.html", u)
	return as.getTokens(u.ID)
}

func (as *authService) loginUser(v *validator.Validator, email, password string) (string, string, error) {
	u, err := as.userSvc.VerifyAndGetUser(v, email, password)
	if err != nil {
		return "", "", err
	}

	return as.getTokens(u.ID)
}

func (as *authService) sendEmail(recipient, templateFile string, u *user.User) {
	go func() {
		err := as.mailer.Send(recipient, templateFile, u)
		if err != nil {
			log.Printf("failed to send email to %s: %v", recipient, err)
		}
	}()
}
