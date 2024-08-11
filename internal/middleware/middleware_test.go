package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bldg14/eventual/internal/middleware"
)

func TestCORS(t *testing.T) {
	tests := []struct {
		name             string
		allowedOrigins   []string
		origin           string
		expectHeader     string
		expectBadOrigins bool
	}{
		{
			name:           "single allowedOrigins",
			allowedOrigins: []string{"https://test.com"},
			origin:         "https://test.com",
			expectHeader:   "https://test.com",
		},
		{
			name:           "multiple allowedOrigins",
			allowedOrigins: []string{"https://test.com", "https://test2.com"},
			origin:         "https://test2.com",
			expectHeader:   "https://test2.com",
		},
		{
			name:           "empty allowedOrigins",
			allowedOrigins: []string{"https://test.com", "", "https://test2.com"},
			origin:         "https://test.com",
			expectHeader:   "https://test.com",
		},
		{
			name:             "no allowedOrigins",
			allowedOrigins:   []string{},
			expectBadOrigins: true,
		},
		{
			name:             "nil allowedOrigins",
			allowedOrigins:   nil,
			expectBadOrigins: true,
		},
		{
			name:             "all empty allowedOrigins",
			allowedOrigins:   []string{"", "", ""},
			expectBadOrigins: true,
		},
		{
			name:           "wildcard allowedOrigins",
			allowedOrigins: []string{"*"},
			origin:         "https://test.com",
			expectHeader:   "https://test.com",
		},
		{
			name:           "wildcard with suffix allowedOrigins",
			allowedOrigins: []string{"*.test.com"},
			origin:         "https://sub.test.com",
			expectHeader:   "https://sub.test.com",
		},
		{
			name:           "wildcard with suffix and non-suffixed allowedOrigins",
			allowedOrigins: []string{"*.test2.com", "https://test.com"},
			origin:         "https://test.com",
			expectHeader:   "https://test.com",
		},
		{
			name:           "wildcard with suffix allowedOrigins and bad origin",
			allowedOrigins: []string{"*.test.com"},
			origin:         "https://test2.com",
			expectHeader:   "",
		},
		{
			name:           "wildcard with suffix allowedOrigins and origin not at suffix",
			allowedOrigins: []string{"*.test.com"},
			origin:         "https://test.com",
			expectHeader:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			allowedOrigins, err := middleware.ParseAllowedOrigins(test.allowedOrigins...)
			if test.expectBadOrigins {
				if err != nil {
					return
				}

				t.Fatalf("expected error but got nothing")
			}

			mw := middleware.CORS(allowedOrigins)

			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			w := httptest.NewRecorder()
			r := http.Request{
				Header: http.Header{},
			}
			r.Header.Set("Origin", test.origin)

			handler := mw(h)
			handler.ServeHTTP(w, &r)

			if w.Code != http.StatusOK {
				t.Errorf("expected %d, got %d", http.StatusOK, w.Code)
			}

			if w.Header().Get("Access-Control-Allow-Origin") != test.expectHeader {
				t.Errorf("expected %q, got %q", test.expectHeader, w.Header().Get("Access-Control-Allow-Origin"))
			}

			if w.Header().Get("Access-Control-Allow-Headers") != "Content-Type, Authorization" {
				t.Errorf("expected %q, got %q", "Content-Type, Authorization", w.Header().Get("Access-Control-Allow-Headers"))
			}
		})
	}
}
