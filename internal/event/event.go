package event

import (
	"fmt"
	"time"
)

type Event struct {
	Title       string
	Start       time.Time
	End         time.Time
	Description string
	URL         string
	Email       string
}

type GetAllEvents func() ([]Event, error)

func GetAll(getAllEvents GetAllEvents) ([]Event, error) {
	events, err := getAllEvents()
	if err != nil {
		return nil, fmt.Errorf("failed to getAllEvents: %w", err)
	}

	return events, nil
}
