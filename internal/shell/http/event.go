package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bldg14/eventual/internal/shell/app"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kevinfalting/mux"
)

func HandleGetAllEvents(pool *pgxpool.Pool) mux.ErrHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		events, err := app.GetAllEvents(pool)
		if err != nil {
			return mux.Error(fmt.Errorf("HandleGetAllEvents failed to GetAllEvents: %w", err), http.StatusInternalServerError)
		}

		result, err := json.Marshal(events)
		if err != nil {
			return mux.Error(fmt.Errorf("HandleGetAllEvents failed to Marshal: %w", err), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)

		return nil
	}
}
