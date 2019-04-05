package vendena

import "time"

// The Token model.
type Token struct {
	ID           int64                `json:"id"`
	UUID         string               `json:"uuid"`
	StoreID      int64                `json:"store_id"`
	ClientID     string               `json:"client_id"`
	ClientSecret string               `json:"client_secret"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	Scopes       []AuthorizationScope `json:"scopes"`
	Session      *Session             `json:"-"`
}

// TokenSession represents the constructor.
type TokenSession struct {
	Session
}

// Tokens creates and returns the constructor.
func (api *API) Tokens() TokenSession {
	var s TokenSession
	s.API = api
	s.URI = "tokens"
	s.Options = map[string]string{}
	return s
}

// Find returns a single instance by ID.
func (sess TokenSession) Find(id int64) (object *Token, vendenaError *Error) {
	object = &Token{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess TokenSession) All() (objects []Token, vendenaError *Error) {
	objects = []Token{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess TokenSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess TokenSession) New() Token {
	return Token{
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Token) Save() (vendenaError *Error) {
	_, vendenaError = save(object, *object.Session, object.UUID)
	return
}

// Delete deletes an object.
func (object *Token) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.UUID)
	return
}
