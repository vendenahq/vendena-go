package vendena

import (
	"strconv"
	"time"
)

// The Customer model.
type Customer struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	PasswordConfirm   string    `json:"password_confirm"`
	UseBillingAddress bool      `json:"use_billing_address"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	BillingAddress    *Address  `json:"billing_address"`
	ShippingAddress   *Address  `json:"shipping_address"`
	Session           *Session  `json:"-"`
}

// CustomerSession represents the constructor.
type CustomerSession struct {
	Session
}

// Customers creates and returns the constructor.
func (api *API) Customers() CustomerSession {
	var s CustomerSession
	s.API = api
	s.URI = "customers"
	s.Options = map[string]string{}
	return s
}

// Page sets the paging option.
func (sess CustomerSession) Page(page int) CustomerSession {
	sess.Options["page"] = strconv.Itoa(page)
	return sess
}

// Limit sets the limit option.
func (sess CustomerSession) Limit(limit int) CustomerSession {
	sess.Options["limit"] = strconv.Itoa(limit)
	return sess
}

// Find returns a single instance by ID.
func (sess CustomerSession) Find(id int64) (object *Customer, vendenaError *Error) {
	object = &Customer{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess CustomerSession) All() (objects []Customer, vendenaError *Error) {
	objects = []Customer{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess CustomerSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess CustomerSession) New() Customer {
	return Customer{
		Session: &sess.Session,
	}
}
