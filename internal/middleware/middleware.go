package middleware

import (
	"net/http"
	"strings"

	"github.com/kevinfalting/mux"
)

// CORS is a middleware that adds the Access-Control-Allow-Origin header to
// responses. The allowedOrigins argument must not be empty.
func CORS(allowedOrigins ...string) mux.Middleware {
	filteredOrigins := make([]string, 0, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		if origin != "" {
			filteredOrigins = append(filteredOrigins, origin)
		}
	}

	if len(filteredOrigins) == 0 {
		panic("allowedOrigins must not be empty")
	}

	origins := strings.Join(filteredOrigins, ", ")

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", origins)
			h.ServeHTTP(w, r)
		})
	}
}
