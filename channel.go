package vendena

import (
	"strconv"
	"time"
)

// The Channel model.
type Channel struct {
	ID                int64     `json:"id"`
	Title             string    `json:"title"`
	Enabled           bool      `json:"enabled"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DefaultCurrencyID int       `json:"default_currency_id"`
	DefaultCurrency   *Currency `json:"default_currency"`
	DefaultLocaleID   int       `json:"default_locale_id"`
	DefaultLocale     *Locale   `json:"default_locale"`
	Domains           []Domain  `json:"domains"`
	Session           *Session  `json:"-"`
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
func (sess ChannelSession) Find(id int64) (object *Channel, err error) {
	object = &Channel{}
	_, err = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess ChannelSession) All() (objects []Channel, err error) {
	objects = []Channel{}
	_, err = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess ChannelSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
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
func (object *Channel) Save() (err error) {
	_, err = save(&object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *Channel) Delete() (err error) {
	_, err = delete(*object.Session, object.ID)
	return
}
