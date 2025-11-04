package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/golang-jwt/jwt/v4"
)

func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// handle preflight OPTIONS
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}

	return fn
}

type ContextKey string

const UserIDKey ContextKey = "userID"

func RequireAuthentication(next http.HandlerFunc) http.HandlerFunc {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			custom_errors.RequireAuthenticationErrorResponse(w)
			return
		}

		claims := &token.Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		issuer, userID := claims.Issuer, claims.UserID
		if issuer != os.Getenv("JWT_ISSUER") {
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return fn
}
