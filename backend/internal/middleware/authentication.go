package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/golang-jwt/jwt/v4"
)

type ContextKey string

const (
	CtxUserKey     ContextKey = "userID"
	CtxTokenString ContextKey = "tokenString"
)

func RequireAccess(next http.HandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(CtxUserKey).(*user.User)
		if !ok {
			panic("user missing from context")
		}

		if u.IsAnonymousUser() {
			customerrors.RequireAuthenticationErrorResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	}
	return Authenticate(fn)
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ctx := context.WithValue(r.Context(), CtxUserKey, user.AnonymousUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		headParts := strings.Split(authHeader, " ")
		if len(headParts) != 2 || headParts[0] != "Bearer" {
			customerrors.RequireAuthenticationErrorResponse(w)
			return
		}
		tokenString := headParts[1]
		token, err := token.ValidateJWT(tokenString, jwtSecret)
		if err != nil {
			customerrors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		issuer, ok := claims["iss"].(string)
		if !ok || issuer != os.Getenv("JWT_ISSUER") {
			customerrors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		subStr, ok := claims["sub"].(string)
		if !ok || subStr == "" {
			customerrors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}
		subInt, err := strconv.Atoi(subStr)
		if err != nil {
			customerrors.InvalidAuthenticationTokenErrorResponse(w)
			return
		}

		u := &user.User{
			ID: subInt,
		}
		ctx := context.WithValue(r.Context(), CtxUserKey, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return fn
}

func RequireRefresh(DB *sql.DB, next http.HandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("REFRESH")
		if err != nil {
			customerrors.RequireAuthenticationErrorResponse(w)
			return
		}
		tokenString := cookie.Value
		if tokenString == "" {
			customerrors.RequireAuthenticationErrorResponse(w)
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
			case errors.Is(err, customerrors.ErrNoRecord):
				customerrors.InvalidAuthenticationTokenErrorResponse(w)
			default:
				customerrors.ServerErrorResponse(w, err)
			}
			return
		}
		// Here we *know* the user, since the refresh token row includes user_id
		u := &user.User{
			ID: tk.UserID,
		}
		ctx := context.WithValue(r.Context(), CtxUserKey, u)
		ctx = context.WithValue(ctx, CtxTokenString, tk.TokenString)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
