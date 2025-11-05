package utilities

import (
	"net/http"
	"os"
	"time"
)

func SetTokenCookie(w http.ResponseWriter, tokenName, token, path string, ttl time.Duration) error {
	secure := os.Getenv("COOKIE_SECURE") != "false" // default true

	cookie := http.Cookie{
		Name:     tokenName,
		Value:    token,
		Expires:  time.Now().Add(ttl),
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		Path:     path,
	}

	err := cookie.Valid()
	if err != nil {
		return err
	}
	http.SetCookie(w, &cookie)
	return nil
}
