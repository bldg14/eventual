package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/kevinfalting/mux"
)

// CORS is a middleware that adds the appropriate headers to responses. The
// allowedOrigins argument must not be empty. Adds Access-Control-Allow-Origin
// and Access-Control-Allow-Headers. The rest of CORS is handled by
// github.com/kevinfalting/mux
//
// If a "*" is supplied in the allowed origins, it will override and allow
// everything. You can supply as many specific origins as you'd like, but they
// must include the full url, including the scheme. If you need a wildcarded
// subdomain, prepend the allowed origin with "*." but do not include a scheme.
// Schemes for wildcarded subdomains are not supported (yet).
func CORS(allowedOrigins AllowedOrigins) mux.Middleware {
	isOriginAllowed := func(origin string) bool {
		if _, ok := allowedOrigins.origins["*"]; ok {
			return true
		}

		if _, ok := allowedOrigins.origins[origin]; ok {
			return true
		}

		for allowedOrigin := range allowedOrigins.origins {
			if !strings.HasPrefix(allowedOrigin, "*.") {
				continue
			}

			if strings.HasSuffix(origin, allowedOrigin[1:]) {
				return true
			}
		}

		return false
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if isOriginAllowed(origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			h.ServeHTTP(w, r)
		})
	}
}

type AllowedOrigins struct {
	origins map[string]struct{}
}

func ParseAllowedOrigins(allowedOrigins ...string) (AllowedOrigins, error) {
	origins := map[string]struct{}{}
	for _, origin := range allowedOrigins {
		if origin == "" {
			continue
		}

		origins[origin] = struct{}{}
	}

	if len(origins) == 0 {
		return AllowedOrigins{}, errors.New("allowedOrigins must not be empty")
	}

	return AllowedOrigins{origins: origins}, nil
}
