package data

import (
	"time"
)

// NOTE: The third column is for controlling how the key
// appears in JSON-encoded format
// - for not including, omitempty for showing only if not empty
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
