package vendena

import "time"

// The Store model.
type Store struct {
	ID        int64     `json:"id"`
	AddressID int64     `json:"-"`
	Title     string    `json:"title"`
	VATNumber string    `json:"vat_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Address   *Address  `json:"address"`
	Session   *Session  `json:"-"`
}

// StoreSession represents the constructor.
type StoreSession struct {
	Session
}

// Stores creates and returns the constructor.
func (api *API) Stores() StoreSession {
	var s StoreSession
	s.API = api
	s.URI = "stores"
	s.Options = map[string]string{}
	return s
}

// Find returns a single instance by ID.
func (sess StoreSession) Find(id int64) (object *Store, err error) {
	object = &Store{}
	_, err = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess StoreSession) All() (objects []Store, err error) {
	objects = []Store{}
	_, err = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess StoreSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess StoreSession) New() Store {
	return Store{Session: &sess.Session}
}

// Save creates or updates an object.
func (object *Store) Save() (err error) {
	_, err = save(object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *Store) Delete() (err error) {
	_, err = delete(*object.Session, object.ID)
	return
}
