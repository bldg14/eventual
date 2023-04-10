package middleware

import (
	"net/http"
	"strings"

	"github.com/kevinfalting/mux"
)

// CORS is a middleware that adds the Access-Control-Allow-Origin header to
// responses. The allowedOrigins argument must not be empty.
func CORS(allowedOrigins ...string) mux.Middleware {
	if len(allowedOrigins) == 0 {
		panic("allowedOrigins must not be empty")
	}

	origins := strings.Join(allowedOrigins, ", ")

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", origins)
			h.ServeHTTP(w, r)
		})
	}
}
