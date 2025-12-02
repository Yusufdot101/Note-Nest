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

func (as *authService) registerUser(v *validator.Validator, name, email, password string) (string, string, error) {
	u, err := as.userSvc.NewUser(v, name, email, password)
	if err != nil {
		return "", "", err
	}

	refresh, access, err := as.tokenSvc.GetAccessAndRefreshTokens(u.ID)
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

	return as.tokenSvc.GetAccessAndRefreshTokens(u.ID)
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

		baseURL := os.Getenv("FRONTEND_BASE_URL")
		if baseURL == "" {
			log.Println("FRONTEND_BASE_URL not set")
			return
		}

		data := struct {
			Name       string
			ResetURL   string
			ExpiryTime string
		}{
			Name:       u.Name,
			ResetURL:   fmt.Sprintf("%s/reset-password?token=%s", baseURL, resetToken),
			ExpiryTime: formattedTime,
		}
		as.sendEmail(u.Email, "password_reset.tmpl.html", data)
	}()

	return nil
}

func (as *authService) resetPassword(v *validator.Validator, token, newPassword string) error {
	return as.userSvc.UpdatePasswordUsingToken(v, token, newPassword)
}

func (as *authService) sendEmail(recipient, templateFile string, data any) {
	go func() {
		err := as.mailer.Send(recipient, templateFile, data)
		if err != nil {
			log.Printf("failed to send email to %s: %v", recipient, err)
		}
	}()
}
