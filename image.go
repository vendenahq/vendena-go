package vendena

import "time"

// The Image model.
type Image struct {
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Session   *Session  `json:"-"`
}

// ImageSession represents the constructor.
type ImageSession struct {
	Session
}

// Images creates and returns the constructor.
func (api *API) Images() ImageSession {
	var s ImageSession
	s.API = api
	s.URI = "images"
	s.Options = map[string]string{}
	return s
}

// Find returns a single instance by ID.
func (sess ImageSession) Find(id string) (object *Image, vendenaError *Error) {
	object = &Image{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess ImageSession) All() (objects []Image, vendenaError *Error) {
	objects = []Image{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess ImageSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess ImageSession) New() Image {
	return Image{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Image) Save() (vendenaError *Error) {
	_, vendenaError = save(object, *object.Session, object.UUID)
	return
}

// Delete deletes an object.
func (object *Image) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.UUID)
	return
}
