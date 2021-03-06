package vendena

import "time"

// The PaymentGateway model.
type PaymentGateway struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Session   *Session  `json:"-"`
}

// PaymentGatewaySession represents the constructor.
type PaymentGatewaySession struct {
	Session
}

// PaymentGateways creates and returns the constructor.
func (api *API) PaymentGateways() PaymentGatewaySession {
	var s PaymentGatewaySession
	s.API = api
	s.URI = "payment_gateways"
	s.Options = map[string]string{}
	return s
}

// Find returns a single instance by ID.
func (sess PaymentGatewaySession) Find(id int64) (object *PaymentGateway, vendenaError *Error) {
	object = &PaymentGateway{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess PaymentGatewaySession) All() (objects []PaymentGateway, vendenaError *Error) {
	objects = []PaymentGateway{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess PaymentGatewaySession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}
