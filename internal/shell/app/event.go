package app

import (
	"fmt"
	"time"

	"github.com/bldg14/eventual/internal/shell/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Event struct {
	Title       string `json:"title"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Email       string `json:"email"`
}

func GetAllEvents(pool *pgxpool.Pool) ([]Event, error) {
	storageEvents, err := storage.GetAll(pool)
	if err != nil {
		return nil, fmt.Errorf("HandleGetAllEvents failed to GetAll: %w", err)
	}

	events := make([]Event, len(storageEvents))
	for i, storageEvent := range storageEvents {
		events[i] = Event{
			Title:       storageEvent.Title,
			Start:       storageEvent.Start.Format(time.RFC3339),
			End:         storageEvent.End.Format(time.RFC3339),
			Description: storageEvent.Description,
			URL:         storageEvent.URL,
			Email:       storageEvent.Email,
		}
	}

	return events, nil
}
