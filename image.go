package vendena

import "time"

// The Image model.
type Image struct {
	ID        int64     `json:"id"`
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
func (sess ImageSession) Find(id int64) (object *Image, err error) {
	object = &Image{}
	_, err = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess ImageSession) All() (objects []Image, err error) {
	objects = []Image{}
	_, err = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess ImageSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
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
func (object *Image) Save() (err error) {
	_, err = save(object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *Image) Delete() (err error) {
	_, err = delete(*object.Session, object.ID)
	return
}
