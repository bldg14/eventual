package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bldg14/eventual/internal/event"
	"github.com/kevinfalting/mux"
)

func GetAllEvents(getAllFunc event.GetAllEvents) mux.ErrHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		events, err := event.GetAll(getAllFunc)
		if err != nil {
			return mux.Error(fmt.Errorf("HandleGetAllEvents failed to GetAll: %w", err), http.StatusInternalServerError)
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
