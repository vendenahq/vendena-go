package vendena

import (
	"strconv"
	"time"
)

// The Bid model.
type Bid struct {
	ID         int64     `json:"id"`
	ParentID   int64     `json:"parent_id"`
	AuctionID  int64     `json:"auction_id"`
	CustomerID int64     `json:"customer_id"`
	Amount     float64   `json:"amount"`
	IsProxy    bool      `json:"is_proxy"`
	Enabled    bool      `json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Customer   *Customer `json:"customer"`
	Session    *Session  `json:"-"`
}

// BidSession represents the constructor.
type BidSession struct {
	Session
}

// Bids creates and returns the constructor.
func (api *API) Bids() BidSession {
	var s BidSession
	s.API = api
	s.URI = "bids"
	s.Options = map[string]string{}
	return s
}

// New creates a new empty object.
func (sess BidSession) New() Bid {
	return Bid{Session: &sess.Session}
}

// Save creates or updates an object.
func (object *Bid) Save() (vendenaError *Error) {
	object.Session.URIPrefix = "auctions/" + strconv.FormatInt(object.AuctionID, 10)
	_, vendenaError = save(object, *object.Session, object.ID)
	return
}
