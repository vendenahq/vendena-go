package vendena

import "time"

// The Variant model.
type Variant struct {
	ID                    int64      `json:"id"`
	ProductID             int64      `json:"product_id"`
	PriceModifierID       int64      `json:"price_modifier_id"`
	ProductOptionValueIDs []int64    `json:"option_value_ids"`
	LocationID            int64      `json:"location_id"`
	Title                 string     `json:"title"`
	Price                 float64    `json:"price"`
	TotalPrice            float64    `json:"total_price"`
	Weight                float64    `json:"weight"`
	SKU                   string     `json:"sku"`
	StockLevel            int        `json:"stock_level"`
	StockLevelLowWarning  int        `json:"stock_level_low_warning"`
	SaleStartAt           *time.Time `json:"sale_start_at,omitempty"`
	SaleEndAt             *time.Time `json:"sale_end_at,omitempty"`
	CustomData            string     `json:"custom_data"`
	Enabled               bool       `json:"enabled"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	Location              *Location  `json:"location"`
	Images                []Image    `json:"images"`
	Session               *Session   `json:"-"`
}

// VariantSession represents the constructor.
type VariantSession struct {
	Session
}

// Variants creates and returns the constructor.
func (api *API) Variants() VariantSession {
	var s VariantSession
	return s
}

// New creates a new empty object.
func (sess VariantSession) New() Variant {
	return Variant{
		Enabled: true,
		Session: &sess.Session,
	}
}

// AddImage adds a new Image.
func (object *Variant) AddImage(image Image) {
	object.Images = append(object.Images, image)
}
