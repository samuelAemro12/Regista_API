package models

import "time"

// User is a minimal placeholder for the future user and workspace system.
type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
