package vendena

import "time"

// The LineItem model.
type LineItem struct {
	ID                    int64     `json:"id"`
	ProductID             int64     `json:"product_id"`
	VariantID             int64     `json:"variant_id"`
	ProductOptionValueIDs []int64   `json:"option_value_ids"`
	Quantity              int64     `json:"quantity"`
	Price                 float64   `json:"price"`
	TotalPrice            float64   `json:"total_price"`
	Weight                float64   `json:"weight"`
	Title                 string    `json:"title"`
	SKU                   string    `json:"sku"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Product               *Product  `json:"product"`
	Variant               *Variant  `json:"variant"`
	Session               *Session  `json:"-"`
}

// LineItemSession represents the constructor.
type LineItemSession struct {
	Session
}

// LineItems creates and returns the constructor.
func (api *API) LineItems() LineItemSession {
	var s LineItemSession
	s.API = api
	s.URI = "line_items"
	s.Options = map[string]string{}
	return s
}

// New creates a new empty object.
func (sess LineItemSession) New() LineItem {
	return LineItem{Session: &sess.Session}
}
