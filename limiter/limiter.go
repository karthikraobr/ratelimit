package limiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	l          sync.RWMutex
	userLimits map[string]*rate.Limiter
	limit      rate.Limit
	burst      int
}

func NewRateLimiter(limit time.Duration, burst int) *RateLimiter {
	return &RateLimiter{
		userLimits: make(map[string]*rate.Limiter),
		limit:      rate.Every(limit),
		burst:      burst,
	}

}

func (r *RateLimiter) Add(ip string) *rate.Limiter {
	r.l.Lock()
	defer r.l.Unlock()
	limiter := rate.NewLimiter(r.limit, r.burst)
	r.userLimits[ip] = limiter
	return limiter
}

func (r *RateLimiter) Get(key string) *rate.Limiter {
	r.l.Lock()
	limiter, ok := r.userLimits[key]
	if !ok {
		r.l.Unlock()
		return r.Add(key)
	}
	r.l.Unlock()
	return limiter
}
