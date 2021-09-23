package main

import (
	"net/http"
	"rateLimit/limiter"
	"time"
)

func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main() {
	http.Handle("/profile/reset", limiter.NewRateLimiterMiddleware(24*time.Hour, 3, resetPasswordHandler))
	http.ListenAndServe(":8081", nil)
}
