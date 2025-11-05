package utilities

import (
	"errors"
	"net/http"
	"os"
	"time"
)

func SetJWTCookie(w http.ResponseWriter, tokenType, token, path string) error {
	var expirationTime time.Duration
	var err error
	switch tokenType {
	case "REFRESH":
		expirationTime, err = time.ParseDuration(os.Getenv("REFRESH_JWT_EXPIRATION_TIME"))
	case "ACCESS":
		expirationTime, err = time.ParseDuration(os.Getenv("ACCESS_JWT_EXPIRATION_TIME"))
	default:
		err = errors.New("invalid tokenType")
	}

	if err != nil {
		return err
	}

	secure := os.Getenv("COOKIE_SECURE") != "false" // default true

	cookie := http.Cookie{
		Name:     tokenType,
		Value:    token,
		Expires:  time.Now().Add(expirationTime),
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		Path:     path,
	}

	err = cookie.Valid()
	if err != nil {
		return err
	}
	http.SetCookie(w, &cookie)
	return nil
}
