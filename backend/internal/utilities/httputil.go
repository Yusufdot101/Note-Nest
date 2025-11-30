package utilities

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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

func DeleteTokenCookie(w http.ResponseWriter, tokenName, path string) error {
	secure := os.Getenv("COOKIE_SECURE") != "false" // default true
	cookie := http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now().Add(-24 * time.Hour),
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

func HumanReadableDuration(d time.Duration) string {
	totalSeconds := int(d.Seconds())

	days := totalSeconds / 86400
	totalSeconds %= 86400

	hours := totalSeconds / 3600
	totalSeconds %= 3600

	minutes := totalSeconds / 60

	var parts []string

	if days > 0 {
		if days == 1 {
			parts = append(parts, "1 day")
		} else {
			parts = append(parts, fmt.Sprintf("%d days", days))
		}
	}

	if hours > 0 {
		if hours == 1 {
			parts = append(parts, "1 hour")
		} else {
			parts = append(parts, fmt.Sprintf("%d hours", hours))
		}
	}

	if minutes > 0 {
		if minutes == 1 {
			parts = append(parts, "1 minute")
		} else {
			parts = append(parts, fmt.Sprintf("%d minutes", minutes))
		}
	}

	if len(parts) == 0 {
		return "0 minutes"
	}

	return strings.Join(parts, " ")
}
