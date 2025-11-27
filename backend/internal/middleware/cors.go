package middleware

import (
	"net/http"
	"os"
	"strings"
)

func EnableCORS(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		trustedOrigins := strings.SplitSeq(os.Getenv("TRUSTED_ORIGINS"), ",")
		isTrusted := false
		for trustedOrigin := range trustedOrigins {
			if origin == trustedOrigin {
				isTrusted = true
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}
		}

		// handle preflight OPTIONS
		if r.Method == http.MethodOptions && isTrusted {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
