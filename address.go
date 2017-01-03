package vendena

import "time"

// The Address model.
type Address struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Address1  string    `json:"address1"`
	Address2  string    `json:"address2"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	Postcode  string    `json:"postcode"`
	Phone     string    `json:"phone"`
	Company   string    `json:"company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Session   *Session  `json:"-"`
}

// AddressSession represents the constructor.
type AddressSession struct {
	Session
}

// Addresses creates and returns the constructor.
func (api *API) Addresses() AddressSession {
	var s AddressSession
	s.API = api
	s.URI = "addresses"
	s.Options = map[string]string{}
	return s
}

// Find returns a single instance by ID.
func (sess AddressSession) Find(id int64) (object *Address, err error) {
	object = &Address{}
	_, err = findOne(object, sess.Session, id)
	return
}

// All returns all instances.
func (sess AddressSession) All() (objects []Address, err error) {
	objects = []Address{}
	_, err = findAll(&objects, sess.Session)
	return
}

// Count returns the number of instances.
func (sess AddressSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess AddressSession) New() Address {
	return Address{Session: &sess.Session}
}

// Save creates or updates an object.
func (object *Address) Save() (err error) {
	_, err = save(object, *object.Session, object.ID)
	return
}
