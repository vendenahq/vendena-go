package vendena

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// The Auction model.
type Auction struct {
	ID                 int64     `json:"id"`
	ProductID          int64     `json:"product_id"`
	ReservePrice       float64   `json:"reserve_price"`
	BidIncrementAmount float64   `json:"bid_increment_amount"`
	BuyItNowPrice      float64   `json:"buy_it_now_price"`
	HasBuyItNow        bool      `json:"has_buy_it_now"`
	HasProxyBids       bool      `json:"has_proxy_bids"`
	StartAt            time.Time `json:"start_at,omitempty"`
	EndAt              time.Time `json:"end_at,omitempty"`
	CustomData         string    `json:"custom_data"`
	Enabled            bool      `json:"enabled"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Product            *Product  `json:"product"`
	Session            *Session  `json:"-"`
}

// AuctionSession represents the constructor.
type AuctionSession struct {
	Session
}

// Auctions creates and returns the constructor.
func (api *API) Auctions() AuctionSession {
	var s AuctionSession
	s.API = api
	s.URI = "auctions"
	s.Options = map[string]string{}
	return s
}

// Page sets the paging option.
func (sess AuctionSession) Page(page int) AuctionSession {
	sess.Options["page"] = strconv.Itoa(page)
	return sess
}

// Limit sets the limit option.
func (sess AuctionSession) Limit(limit int) AuctionSession {
	sess.Options["limit"] = strconv.Itoa(limit)
	return sess
}

// Find returns a single instance by ID.
func (sess AuctionSession) Find(id int64) (object *Auction, vendenaError *Error) {
	object = &Auction{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess AuctionSession) All() (objects []Auction, vendenaError *Error) {
	objects = []Auction{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess AuctionSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess AuctionSession) New() Auction {
	return Auction{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Auction) Save() (vendenaError *Error) {
	_, vendenaError = save(object, *object.Session, object.ID)
	return
}

// Bids gets the delivery methods available to the order.
func (object *Auction) DeliveryMethods() (bids []Bid, vendenaError *Error) {
	result, status, vendenaError := request(*object.Session, http.MethodGet, strconv.FormatInt(object.ID, 10), "bids", nil)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&bids); err != nil {
		vendenaError = createError("json_decoder_error", err)
		return
	}

	return
}

// Delete deletes an object.
func (object *Auction) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.ID)
	return
}
