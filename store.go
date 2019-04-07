package vendena

import "time"

// The Store model.
type Store struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	ProjectID int64     `json:"project_id"`
	AddressID int64     `json:"-"`
	Title     string    `json:"title"`
	Currency  string    `json:"currency"`
	Locale    string    `json:"locale"`
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
func (sess StoreSession) Find(id string) (object *Store, vendenaError *Error) {
	object = &Store{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess StoreSession) All() (objects []Store, vendenaError *Error) {
	objects = []Store{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess StoreSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess StoreSession) New() Store {
	return Store{Session: &sess.Session}
}

// Save creates or updates an object.
func (object *Store) Save() (vendenaError *Error) {
	_, vendenaError = save(object, *object.Session, object.UUID)
	return
}

// Delete deletes an object.
func (object *Store) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.UUID)
	return
}
