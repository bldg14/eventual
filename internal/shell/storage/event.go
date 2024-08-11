package storage

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Event struct {
	Title       string
	Start       time.Time
	End         time.Time
	Description string
	URL         string
	Email       string
}

func GetAll(pool *pgxpool.Pool) ([]Event, error) {
	return []Event{
		{
			Title:       "Petting Zoo",
			Start:       time.Now().Add(24 * time.Hour),
			End:         time.Now().Add(30 * time.Hour),
			Description: "Come pet a dog, a cow, a monkey, an alligator, and an elephant!",
			URL:         "pettingzoo.com",
			Email:       "jim@pettingzoo.com",
		},
		{
			Title:       "Rock Concerto",
			Start:       time.Now().Add(40 * time.Hour),
			End:         time.Now().Add(45 * time.Hour),
			Description: "Listen to the sound of rocks",
			URL:         "soundofrocks.com",
			Email:       "jim@soundofrocks.com",
		},
		{
			Title:       "Recycling",
			Start:       time.Now(),
			End:         time.Now().Add(2 * time.Hour),
			Description: "Bring gloves, we're recycling things.",
			URL:         "recycle.com",
			Email:       "jim@recycle.com",
		},
		{
			Title:       "MMA Match",
			Start:       time.Now().Add(40 * time.Hour),
			End:         time.Now().Add(42 * time.Hour),
			Description: "It's a fight, bring you're own popcorn.",
			URL:         "mma.com",
			Email:       "jim@mma.com",
		},
		{
			Title:       "Underwater Basket Weaving",
			Start:       time.Now().Add(72 * time.Hour),
			End:         time.Now().Add(73 * time.Hour),
			Description: "Grab a snorkle and your favorite basket weaving material. Medical standing by.",
			URL:         "deepseaweaving.com",
			Email:       "jim@deepseaweaving.com",
		},
	}, nil
}
