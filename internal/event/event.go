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

type EventStorer interface {
	GetEvents() ([]Event, error)
}

func GetAll(s EventStorer) ([]Event, error) {
	events, err := s.GetEvents()
	if err != nil {
		return nil, fmt.Errorf("failed to GetEvents: %w", err)
	}

	return events, nil
}
