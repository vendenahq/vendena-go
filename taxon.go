package vendena

import (
	"strconv"
	"time"
)

// The Taxon model.
type Taxon struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	ParentID  int64     `json:"parent_id"`
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Session   *Session  `json:"-"`
}

// TaxonSession represents the constructor.
type TaxonSession struct {
	Session
}

// Taxons creates and returns the constructor.
func (api *API) Taxons() TaxonSession {
	var s TaxonSession
	s.API = api
	s.URI = "taxons"
	s.Options = map[string]string{}
	return s
}

// ParentID sets the parent_id option.
func (sess TaxonSession) ParentID(id int64) TaxonSession {
	sess.Options["parent_id"] = strconv.FormatInt(id, 10)
	return sess
}

// SiblingID sets the sibling_id option.
func (sess TaxonSession) SiblingID(id int64) TaxonSession {
	sess.Options["sibling_id"] = strconv.FormatInt(id, 10)
	return sess
}

// Find returns a single instance by ID.
func (sess TaxonSession) Find(id string) (object *Taxon, vendenaError *Error) {
	object = &Taxon{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess TaxonSession) All() (objects []Taxon, vendenaError *Error) {
	objects = []Taxon{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess TaxonSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess TaxonSession) New() Taxon {
	return Taxon{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Taxon) Save() (vendenaError *Error) {
	_, vendenaError = save(object, *object.Session, object.UUID)
	return
}

// Delete deletes an object.
func (object *Taxon) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.UUID)
	return
}
