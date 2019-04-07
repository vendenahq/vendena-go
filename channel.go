package vendena

import (
	"strconv"
	"time"
)

// The Channel model.
type Channel struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Currency  string    `json:"currency"`
	Locale    string    `json:"locale"`
	Domains   []Domain  `json:"domains"`
	Session   *Session  `json:"-"`
}

// ChannelSession represents the constructor.
type ChannelSession struct {
	Session
}

// Channels creates and returns the constructor.
func (api *API) Channels() ChannelSession {
	var s ChannelSession
	s.API = api
	s.URI = "channels"
	s.Options = map[string]string{}
	return s
}

// Page sets the paging option.
func (sess ChannelSession) Page(page int) ChannelSession {
	sess.Options["page"] = strconv.Itoa(page)
	return sess
}

// Limit sets the limit option.
func (sess ChannelSession) Limit(limit int) ChannelSession {
	sess.Options["limit"] = strconv.Itoa(limit)
	return sess
}

// Find returns a single instance by ID.
func (sess ChannelSession) Find(id int64) (object *Channel, vendenaError *Error) {
	object = &Channel{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess ChannelSession) All() (objects []Channel, vendenaError *Error) {
	objects = []Channel{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess ChannelSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess ChannelSession) New() Channel {
	return Channel{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Channel) Save() (vendenaError *Error) {
	_, vendenaError = save(&object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *Channel) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.ID)
	return
}
