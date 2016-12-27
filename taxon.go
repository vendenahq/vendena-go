package vendena

import "time"

// The Taxon model.
type Taxon struct {
	ID        int64     `json:"id"`
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

// Find returns a single instance by ID.
func (sess TaxonSession) Find(id int64) (object *Taxon, err error) {
	object = &Taxon{}
	_, err = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess TaxonSession) All() (objects []Taxon, err error) {
	objects = []Taxon{}
	_, err = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess TaxonSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
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
func (object *Taxon) Save() (err error) {
	_, err = save(object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *Taxon) Delete() (err error) {
	_, err = delete(*object.Session, object.ID)
	return
}
