package vendena

import (
	"strconv"
	"time"
)

// The ProductOption model.
type ProductOption struct {
	ID        int64                `json:"id"`
	Code      string               `json:"code"`
	Title     string               `json:"title"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Values    []ProductOptionValue `json:"values"`
	Session   *Session             `json:"-"`
}

// ProductOptionSession represents the constructor.
type ProductOptionSession struct {
	Session
}

// ProductOptions creates and returns the constructor.
func (api *API) ProductOptions() ProductOptionSession {
	var s ProductOptionSession
	s.API = api
	s.URI = "product_options"
	s.Options = map[string]string{}
	return s
}

// Page sets the paging option.
func (sess ProductOptionSession) Page(page int) ProductOptionSession {
	sess.Options["page"] = strconv.Itoa(page)
	return sess
}

// Limit sets the limit option.
func (sess ProductOptionSession) Limit(limit int) ProductOptionSession {
	sess.Options["limit"] = strconv.Itoa(limit)
	return sess
}

// Find returns a single instance by ID.
func (sess ProductOptionSession) Find(id int64) (object *ProductOption, err error) {
	object = &ProductOption{}
	_, err = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess ProductOptionSession) All() (objects []ProductOption, err error) {
	objects = []ProductOption{}
	_, err = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess ProductOptionSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess ProductOptionSession) New() ProductOption {
	return ProductOption{
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *ProductOption) Save() (err error) {
	_, err = save(object, *object.Session, object.ID)
	return
}

// AddValue adds a new ProductOptionValue.
func (object *ProductOption) AddValue(v ProductOptionValue) {
	object.Values = append(object.Values, v)
	return
}

// Delete deletes an object.
func (object *ProductOption) Delete() (err error) {
	_, err = delete(*object.Session, object.ID)
	return
}
