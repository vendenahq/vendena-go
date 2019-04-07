package vendena

import "time"

// The Domain model.
type Domain struct {
	ID        int64     `json:"id"`
	ChannelID int64     `json:"channel_id"`
	Host      string    `json:"host"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Currency  string    `json:"currency"`
	Locale    string    `json:"locale"`
}
