package utilities

import (
	"net/http"
	"os"
	"time"
)

func SetJWTCookie(w http.ResponseWriter, token string) error {
	expirationTime, err := time.ParseDuration(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		return err
	}

	secure := os.Getenv("COOKIE_SECURE") != "false" // default true

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(expirationTime),
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	err = cookie.Valid()
	if err != nil {
		return err
	}
	http.SetCookie(w, &cookie)
	return nil
}
