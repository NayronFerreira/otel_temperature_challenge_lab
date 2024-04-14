package middleware

import (
	"net/http"
	"strings"

	"github.com/NayronFerreira/microservice-ratelimiter/config"
	limiter "github.com/NayronFerreira/microservice-ratelimiter/ratelimiter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func RateLimitMiddleware(next http.Handler, rateLimiter *limiter.RateLimiter, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		carrier := propagation.HeaderCarrier(r.Header)
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), carrier)
		tracer := otel.Tracer("microservice-ratelimiter")

		ctx, span := tracer.Start(ctx, "check-ratelimit")
		defer span.End()
		token := r.Header.Get("API_KEY")

		if token != "" && rateLimiter.TokenExists(token) {

			isBlocked, err := rateLimiter.CheckRateLimitForKey(ctx, token, true)
			if err != nil {
				http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			if isBlocked {
				http.Error(w, "Your Token have reached the maximum number of requests or actions allowed within a certain time frame.", http.StatusTooManyRequests)
				return
			}

		} else {

			ip := strings.Split(r.RemoteAddr, ":")[0]
			isBlocked, err := rateLimiter.CheckRateLimitForKey(ctx, ip, false)
			if err != nil {
				http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
				return
			}

			if isBlocked {
				http.Error(w, "Your IP have reached the maximum number of requests or actions allowed within a certain time frame.", http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
