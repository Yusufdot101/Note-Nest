package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/golang-jwt/jwt/v4"
)

func EnableCORS(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// handle preflight OPTIONS
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

type ContextKey string

const UserIDKey ContextKey = "userID"

func RequireAuthentication(next http.HandlerFunc, tokenType string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var jwtSecret []byte
		switch tokenType {
		case "REFRESH":
			jwtSecret = []byte(os.Getenv("REFRESH_JWT_SECRET"))
		case "ACCESS":
			jwtSecret = []byte(os.Getenv("ACCESS_JWT_SECRET"))
		default:
			jwtSecret = []byte(os.Getenv(""))
		}
		if len(jwtSecret) == 0 {
			custom_errors.ServerErrorResponse(w, errors.New("JWT_SECRET variable is not set"))
			return
		}

		cookie, err := r.Cookie("jwt")
		if err != nil {
			custom_errors.RequireAuthenticationErrorResponse(w)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (any, error) {
			// ensure the token was signed with HMAC, not something else
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		issuer, userID := claims["iss"], claims["sub"]
		if issuer != os.Getenv("JWT_ISSUER") {
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
