package vendena

import "time"

// The Domain model.
type Domain struct {
	ID              int64     `json:"id"`
	UUID            string    `json:"uuid"`
	ChannelID       int64     `json:"channel_id"`
	Host            string    `json:"host"`
	Enabled         bool      `json:"enabled"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DefaultCurrency *Currency `json:"default_currency"`
	DefaultLocale   *Locale   `json:"default_locale"`
}
