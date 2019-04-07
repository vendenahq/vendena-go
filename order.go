package vendena

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// The Order model.
type Order struct {
	ID                  int64                   `json:"id"`
	ChannelID           int64                   `json:"channel_id"`
	CartID              int64                   `json:"cart_id"`
	CustomerID          int64                   `json:"customer_id"`
	Token               string                  `json:"token"`
	StatusID            int64                   `json:"status_id"`
	Status              *OrderStatus            `json:"status"`
	PaymentStatusID     int64                   `json:"payment_status_id"`
	PaymentStatus       *OrderPaymentStatus     `json:"payment_status"`
	FulfillmentStatusID int64                   `json:"fulfillment_status_id"`
	FulfillmentStatus   *OrderFulfillmentStatus `json:"fulfillment_status"`
	DeliveryMethodID    int64                   `json:"delivery_method_id"`
	DeliveryMethod      *DeliveryMethod         `json:"delivery_method"`
	PaymentMethodID     int64                   `json:"payment_method_id"`
	PaymentMethod       *PaymentMethod          `json:"payment_method"`
	UseBillingAddress   bool                    `json:"use_billing_address"`
	MessageFromBuyer    string                  `json:"message_from_buyer"`
	SubTotal            float64                 `json:"sub_total"`
	TotalPrice          float64                 `json:"total_price"`
	Weight              float64                 `json:"weight"`
	Currency            string                  `json:"currency"`
	Locale              string                  `json:"locale"`
	CreatedAt           time.Time               `json:"created_at"`
	UpdatedAt           time.Time               `json:"updated_at"`
	Customer            *Customer               `json:"customer"`
	BillingAddress      *Address                `json:"billing_address"`
	ShippingAddress     *Address                `json:"shipping_address"`
	Transactions        []Transaction           `json:"transactions"`
	LineItems           []LineItem              `json:"line_items"`
	Revisions           []OrderRevision         `json:"revisions"`
	Session             *Session                `json:"-"`
}

// OrderSession represents the constructor.
type OrderSession struct {
	Session
}

// Orders creates and returns the constructor.
func (api *API) Orders() OrderSession {
	var s OrderSession
	s.API = api
	s.URI = "orders"
	s.Options = map[string]string{}
	return s
}

// FindByToken returns a single instance by token.
func (sess OrderSession) FindByToken(token string) (object *Order, vendenaError *Error) {
	object = &Order{}
	_, vendenaError = findOneByToken(object, sess.Session, token)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess OrderSession) All() (objects []Order, vendenaError *Error) {
	objects = []Order{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess OrderSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess OrderSession) New() Order {
	return Order{
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Order) Save() (vendenaError *Error) {
	_, vendenaError = saveByToken(object, *object.Session, object.Token)
	return
}

// DeliveryMethods gets the delivery methods available to the order.
func (object *Order) DeliveryMethods() (deliveryMethods []DeliveryMethod, vendenaError *Error) {
	result, status, vendenaError := request(*object.Session, http.MethodGet, object.Token, "delivery_methods", nil)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&deliveryMethods); err != nil {
		vendenaError = createError("json_decoder_error", err)
		return
	}

	return
}

// PaymentMethods gets the payment methods available to the order.
func (object *Order) PaymentMethods() (paymentMethods []PaymentMethod, vendenaError *Error) {
	result, status, vendenaError := request(*object.Session, http.MethodGet, object.Token, "payment_methods", nil)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&paymentMethods); err != nil {
		vendenaError = createError("json_decoder_error", err)
		return
	}

	return
}

// FormData gets the order's payment form data for external checkout integrations.
func (object *Order) FormData() (formData OrderFormData, vendenaError *Error) {
	result, status, err := request(*object.Session, http.MethodGet, object.Token, "payment_form", nil)

	if err != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&formData); err != nil {
		vendenaError = createError("json_decoder_error", err)
		return
	}

	if err != nil {
		return
	}

	return
}

// ValidateNotification validates the notification callback from an external checkout integration.
func (object *Order) ValidateNotification(data string) (validationResult OrderNotificationValidationResult, vendenaError *Error) {
	var validation = OrderNotificationValidation{
		Querystring: data,
	}

	var body = &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(&validation); err != nil {
		vendenaError = createError("json_encoder_error", err)
		return
	}

	result, status, vendenaError := request(*object.Session, http.MethodPost, object.Token, "validate_notification", body)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&validationResult); err != nil {
		vendenaError = createError("json_decoder_error", err)
		return
	}

	return
}
