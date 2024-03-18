package models

import "time"

type Event struct {
	ID        int       `json:"event_id" db:"event_id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"start_date" db:"start_date"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
