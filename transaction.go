package vendena

import "time"

// The Transaction model.
type Transaction struct {
	ID                int64     `json:"id"`
	OrderID           int64     `json:"order_id"`
	Amount            float64   `json:"amount"`
	ResponseCode      string    `json:"response_code"`
	AuthorizationCode string    `json:"authorization_code"`
	ReferenceCode     string    `json:"reference_code"`
	TransactionCode   string    `json:"transaction_code"`
	CreditCardType    string    `json:"credit_card_type"`
	CreditCardNumber  string    `json:"credit_card_number"`
	Failed            bool      `json:"failed"`
	FailureReason     string    `json:"failure_reason"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
