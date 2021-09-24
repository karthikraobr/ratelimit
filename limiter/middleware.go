package limiter

import (
	"encoding/json"
	"net/http"
	"time"
)

func NewRateLimiterMiddleware(limit time.Duration, burst int, next http.HandlerFunc) http.Handler {
	rateLimitter := NewRateLimiter(limit, burst)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var email struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		limiter := rateLimitter.Get(email.Email)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next(w, r)
	})
}
