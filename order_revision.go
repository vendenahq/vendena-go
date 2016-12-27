package vendena

import "time"

// The OrderRevision model.
type OrderRevision struct {
	ID                int64                   `json:"id"`
	Notes             string                  `json:"notes"`
	CreatedAt         time.Time               `json:"created_at"`
	UpdatedAt         time.Time               `json:"updated_at"`
	Status            *OrderStatus            `json:"status"`
	PaymentStatus     *OrderPaymentStatus     `json:"payment_status"`
	FulfillmentStatus *OrderFulfillmentStatus `json:"fulfillment_status"`
}
