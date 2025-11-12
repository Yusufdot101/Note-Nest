package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/golang-jwt/jwt/v4"
)

func EnableCORS(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		trustedOrigins := strings.SplitSeq(os.Getenv("TRUSTED_ORIGINS"), ",")
		for trustedOrigin := range trustedOrigins {
			if origin == trustedOrigin {
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}
		}

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

const (
	CtxUserIDKey   ContextKey = "userID"
	CtxTokenString ContextKey = "tokenString"
)

func RequireAccess(next http.HandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		authHeader := r.Header.Get("Authorization")
		headParts := strings.Split(authHeader, " ")
		if len(headParts) != 2 || headParts[0] != "Bearer" {
			custom_errors.RequireAuthenticationErrorResponse(w)
			return
		}
		tokenString := headParts[1]
		token, err := token.ValidateJWT(tokenString, jwtSecret)
		if err != nil {
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		issuer, ok := claims["iss"].(string)
		if !ok || issuer != os.Getenv("JWT_ISSUER") {
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		subStr, ok := claims["sub"].(string)
		if !ok || subStr == "" {
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}
		subInt, err := strconv.Atoi(subStr)
		if err != nil {
			custom_errors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), CtxUserIDKey, subInt)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func RequireRefresh(DB *sql.DB, next http.HandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("REFRESH")
		if err != nil {
			custom_errors.RequireAuthenticationErrorResponse(w)
			return
		}
		tokenString := cookie.Value
		if tokenString == "" {
			custom_errors.RequireAuthenticationErrorResponse(w)
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
		// Here we *know* the user, since the refresh token row includes user_id
		ctx := context.WithValue(r.Context(), CtxUserIDKey, tk.UserID)
		ctx = context.WithValue(ctx, CtxTokenString, tk.TokenString)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
