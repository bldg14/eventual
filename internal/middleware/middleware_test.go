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
		expectHeader   string
		expectPanic    bool
	}{
		{
			name:           "single origin",
			allowedOrigins: []string{"hello, test"},
			expectHeader:   "hello, test",
			expectPanic:    false,
		},
		{
			name:           "multiple origins",
			allowedOrigins: []string{"hello, test", "hello, test2"},
			expectHeader:   "hello, test, hello, test2",
			expectPanic:    false,
		},
		{
			name:           "empty origin",
			allowedOrigins: []string{"hello, test", "", "hello, test2"},
			expectHeader:   "hello, test, hello, test2",
			expectPanic:    false,
		},
		{
			name:           "no origins",
			allowedOrigins: []string{},
			expectPanic:    true,
		},
		{
			name:           "nil origins",
			allowedOrigins: nil,
			expectPanic:    true,
		},
		{
			name:           "all empty origins",
			allowedOrigins: []string{"", "", ""},
			expectPanic:    true,
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
			r := http.Request{}

			handler := mw(h)
			handler.ServeHTTP(w, &r)

			if w.Code != http.StatusOK {
				t.Errorf("expected %d, got %d", http.StatusOK, w.Code)
			}

			if w.Header().Get("Access-Control-Allow-Origin") != test.expectHeader {
				t.Errorf("expected %s, got %s", test.expectHeader, w.Header().Get("Access-Control-Allow-Origin"))
			}

			isPanicking = false
		})
	}
}
