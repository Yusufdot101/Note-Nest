package auth

import (
	"github.com/Yusufdot101/note-nest/internal/token"
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

	return as.getTokens(u.ID)
}

func (as *authService) loginUser(v *validator.Validator, email, password string) (string, string, error) {
	u, err := as.userSvc.VerifyAndGetUser(v, email, password)
	if err != nil {
		return "", "", err
	}

	return as.getTokens(u.ID)
}
