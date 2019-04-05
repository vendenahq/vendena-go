package vendena

import "time"

// The ProductOptionValue model.
type ProductOptionValue struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Session   *Session  `json:"-"`
}

// ProductOptionValueSession represents the constructor.
type ProductOptionValueSession struct {
	Session
}

// ProductOptionValues creates and returns the constructor.
func (api *API) ProductOptionValues() ProductOptionValueSession {
	var s ProductOptionValueSession
	return s
}

// New creates a new empty object.
func (sess ProductOptionValueSession) New() ProductOptionValue {
	return ProductOptionValue{Session: &sess.Session}
}
