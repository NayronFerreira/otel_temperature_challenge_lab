package middleware

import (
	"io"
	"net/http"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req, err := http.NewRequestWithContext(r.Context(), "GET", "http://microservice-ratelimiter:8080/check", nil)
		if err != nil {
			http.Error(w, "Error creating request to rate limiter:"+err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Error contacting rate limiter:"+err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading rate limiter response", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusTooManyRequests {
			http.Error(w, "Too many requests: "+string(body), resp.StatusCode)
			return
		}

		next.ServeHTTP(w, r)
	})
}
