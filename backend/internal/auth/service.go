package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/utilities"
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

	refresh, access, err := as.getTokens(u.ID)
	if err != nil {
		return "", "", err
	}
	as.sendEmail(u.Email, "user_welcome.tmpl.html", u)
	return refresh, access, nil
}

func (as *authService) loginUser(v *validator.Validator, email, password string) (string, string, error) {
	u, err := as.userSvc.VerifyAndGetUser(v, email, password)
	if err != nil {
		return "", "", err
	}

	return as.getTokens(u.ID)
}

func (as *authService) forgotPassword(v *validator.Validator, email string) error {
	u, err := as.userSvc.GetUserByEmail(v, email)
	if err != nil {
		return err
	}
	go func() {
		ttl, err := time.ParseDuration(os.Getenv("RESET_PASSWORD_TOKEN_EXPIRATION_TIME"))
		if err != nil {
			log.Println("error parsing time: ", err)
			return
		}

		formattedTime := utilities.HumanReadableDuration(ttl)

		resetToken, err := as.tokenSvc.NewToken(token.RANDOMSTRING, token.RESET, u.ID)
		if err != nil {
			log.Println("token generation error: ", err)
			return
		}
		data := struct {
			Name       string
			ResetURL   string
			ExpiryTime string
		}{
			Name:       u.Name,
			ResetURL:   fmt.Sprintf("http://localhost:3000/password-reset?%s", resetToken),
			ExpiryTime: formattedTime,
		}
		as.sendEmail(u.Email, "password_reset.tmpl.html", data)
	}()

	return nil
}

func (as *authService) sendEmail(recipient, templateFile string, data any) {
	go func() {
		err := as.mailer.Send(recipient, templateFile, data)
		if err != nil {
			log.Printf("failed to send email to %s: %v", recipient, err)
		}
	}()
}
