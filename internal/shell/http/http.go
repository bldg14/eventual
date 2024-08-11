package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bldg14/eventual/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinfalting/mux"
)

type Config struct {
	Port           int
	AllowedOrigins []string
	Pool           *pgxpool.Pool
}

func NewServer(cfg Config) (*http.Server, error) {
	allowedOrigins, err := middleware.ParseAllowedOrigins(cfg.AllowedOrigins...)
	if err != nil {
		return nil, fmt.Errorf("failed to ParseAllowedOrigins: %w", err)
	}

	api := mux.New(
		middleware.CORS(allowedOrigins),
	)

	eh := mux.ErrorHandler{
		ErrWriter: os.Stderr,
		ErrFunc:   http.Error,
	}

	api.Handle("/api/v1/events", mux.Methods(
		mux.WithGET(eh.Err(HandleGetAllEvents(cfg.Pool))),
	))

	server := http.Server{
		Handler: api,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	return &server, nil
}
