package vendena

import (
	"strconv"
	"time"
)

// The Location model.
type Location struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Address   *Address  `json:"address"`
	Session   *Session  `json:"-"`
}

// LocationSession represents the constructor.
type LocationSession struct {
	Session
}

// Locations creates and returns the constructor.
func (api *API) Locations() LocationSession {
	var s LocationSession
	s.API = api
	s.URI = "locations"
	s.Options = map[string]string{}
	return s
}

// Page sets the paging option.
func (sess LocationSession) Page(page int) LocationSession {
	sess.Options["page"] = strconv.Itoa(page)
	return sess
}

// Limit sets the limit option.
func (sess LocationSession) Limit(limit int) LocationSession {
	sess.Options["limit"] = strconv.Itoa(limit)
	return sess
}

// Find returns a single instance by ID.
func (sess LocationSession) Find(id int64) (object *Location, vendenaError *Error) {
	object = &Location{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess LocationSession) All() (objects []Location, vendenaError *Error) {
	objects = []Location{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess LocationSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess LocationSession) New() Location {
	return Location{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Location) Save() (vendenaError *Error) {
	_, vendenaError = save(&object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *Location) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.ID)
	return
}
