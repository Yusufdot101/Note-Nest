package middleware

import (
	"fmt"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
)

func RecoverPanic(next http.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				customerrors.ServerErrorResponse(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	}
	return fn
}
