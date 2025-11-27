package middleware

import (
	"net/http"
	"sync"
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

	var mu sync.Mutex
	clients := make(map[string]*client)

	go func() {
		for {
			time.Sleep(1 * time.Minute)
			mu.Lock()
			for ip, c := range clients {
				// delete the limiter if the client didnt use the api in the last 3 minutes
				if time.Since(c.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	fn := func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiterEnabled {
			next.ServeHTTP(w, r)
			return
		}

		ip := realip.FromRequest(r)
		mu.Lock()
		if _, exists := clients[ip]; !exists {
			clients[ip] = &client{
				limiter: rate.NewLimiter(rate.Limit(rateLimitRate), rateLimitBurst),
			}
		}
		clients[ip].lastSeen = time.Now()
		mu.Unlock()

		if !clients[ip].limiter.Allow() {
			customerrors.RateLimitExceededErrorResponse(w)
			return
		}
		next.ServeHTTP(w, r)
	}

	return fn
}
