package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func Logger(l *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Create a new logger for this request
			reqlog := l.With().
				Str("method", r.Method).
				Str("url", r.URL.RequestURI()).
				Str("remote_addr", r.RemoteAddr).
				Logger()

			// Create a response writer that captures the status code
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			// Log the request start
			reqlog.Info().Msg("request started")

			// Start a timer
			t1 := time.Now()

			defer func() {
				// Log the request completion
				reqlog.Info().
					Int("status", ww.Status()).
					Int("bytes", ww.BytesWritten()).
					Dur("duration", time.Since(t1)).
					Msg("request completed")
			}()

			// Call the next handler in the chain
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
