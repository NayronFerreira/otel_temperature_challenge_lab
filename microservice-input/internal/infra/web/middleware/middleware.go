package middleware

import (
	"io"
	"net/http"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		carrier := propagation.HeaderCarrier(r.Header)
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), carrier)
		tracer := otel.Tracer(viper.GetString("SERVICE_NAME"))

		ctx, span := tracer.Start(ctx, "web-handler")
		defer span.End()

		req, err := http.NewRequestWithContext(ctx, "GET", "http://microservice-ratelimiter:8080/check", nil)
		if err != nil {
			http.Error(w, "Error creating request to rate limiter:"+err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Header.Get("API_KEY") != "" {
			req.Header.Set("API_KEY", r.Header.Get("API_KEY"))
		}

		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

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

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
