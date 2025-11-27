package middleware

import (
	"net/http"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

func RateLimit(next http.Handler, rateLimiterEnabled bool, rateLimitBurst int, rateLimitRate float64) http.HandlerFunc {
	type client struct {
		lastSeen time.Time
		limiter  *rate.Limiter
	}

	clients := make(map[string]*client)
	go func() {
		for {
			for ip, c := range clients {
				// delete the limiter if the client didnt use the api in the last 3 minutes
				if time.Since(c.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
		}
	}()

	fn := func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiterEnabled {
			next.ServeHTTP(w, r)
			return
		}

		ip := realip.FromRequest(r)
		if _, exists := clients[ip]; !exists {
			clients[ip] = &client{
				limiter: rate.NewLimiter(rate.Limit(rateLimitBurst), rateLimitBurst),
			}
		}
		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			customerrors.RateLimitExceededErrorResponse(w)
			return
		}
		next.ServeHTTP(w, r)
	}

	return fn
}
