package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bldg14/eventual/internal/middleware"
)

func TestCORS(t *testing.T) {
	tests := []struct {
		name           string
		allowedOrigins []string
		origin         string
		expectHeader   string
		expectPanic    bool
	}{
		{
			name:           "single allowedOrigins",
			allowedOrigins: []string{"https://test.com"},
			origin:         "https://test.com",
			expectHeader:   "https://test.com",
			expectPanic:    false,
		},
		{
			name:           "multiple allowedOrigins",
			allowedOrigins: []string{"https://test.com", "https://test2.com"},
			origin:         "https://test2.com",
			expectHeader:   "https://test2.com",
			expectPanic:    false,
		},
		{
			name:           "empty allowedOrigins",
			allowedOrigins: []string{"https://test.com", "", "https://test2.com"},
			origin:         "https://test.com",
			expectHeader:   "https://test.com",
			expectPanic:    false,
		},
		{
			name:           "no allowedOrigins",
			allowedOrigins: []string{},
			expectPanic:    true,
		},
		{
			name:           "nil allowedOrigins",
			allowedOrigins: nil,
			expectPanic:    true,
		},
		{
			name:           "all empty allowedOrigins",
			allowedOrigins: []string{"", "", ""},
			expectPanic:    true,
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
			isPanicking := true
			defer func() {
				r := recover()
				if isPanicking && !test.expectPanic {
					t.Errorf("expected no panic, got %v", r)
				}
			}()

			mw := middleware.CORS(test.allowedOrigins...)

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

			isPanicking = false
		})
	}
}
