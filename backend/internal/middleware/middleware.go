package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/token"
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

func RequireAuthentication(next http.HandlerFunc, tokenUse token.TokenUse, DB *sql.DB) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var userID int
		switch tokenUse {
		case token.ACCESS:
			jwtSecret := []byte(os.Getenv("JWT_SECRET"))
			authHeader := r.Header.Get("Authorization")
			headParts := strings.Split(authHeader, " ")
			if len(headParts) != 2 || headParts[0] != "Bearer" {
				custom_errors.InvalidAuthenticationTokenErrorResponse(w)
				return
			}
			tokenString := headParts[1]

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
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
			issuer, ok := claims["iss"].(string)
			if !ok || issuer != os.Getenv("JWT_ISSUER") {
				custom_errors.InvalidAuthenticationTokenErrorResponse(w)
				return
			}

			sub, ok := claims["sub"].(string)
			if !ok || sub == "" {
				custom_errors.InvalidAuthenticationTokenErrorResponse(w)
				return
			}
			subInt, ok := claims["sub"].(int)
			if !ok || sub == "" {
				custom_errors.InvalidAuthenticationTokenErrorResponse(w)
				return
			}
			userID = subInt

		case token.REFRESH:
			cookie, err := r.Cookie("REFRESH")
			if err != nil {
				custom_errors.InvalidAuthenticationTokenErrorResponse(w)
				return
			}
			tokenString := cookie.Value
			if tokenString == "" {
				custom_errors.InvalidAuthenticationTokenErrorResponse(w)
				return
			}
			svc := &token.TokenService{
				Repo: &token.Repository{
					DB: DB,
				},
			}
			tk, err := svc.Repo.GetByTokenString(tokenString)
			if err != nil {
				switch {
				case errors.Is(err, custom_errors.ErrNoRecord):
					custom_errors.InvalidAuthenticationTokenErrorResponse(w)
				default:
					custom_errors.ServerErrorResponse(w, err)
				}
				return
			}
			userID = tk.UserID
		default:
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
