package vendena

import (
	"strconv"
	"time"
)

// The DeliveryMethod model.
type DeliveryMethod struct {
	ID                   int64               `json:"id"`
	DeliveryMethodTypeID int64               `json:"delivery_method_type_id"`
	Title                string              `json:"title"`
	Details              string              `json:"details"`
	Instructions         string              `json:"instructions"`
	RangeFrom            float64             `json:"range_from"`
	RangeTo              float64             `json:"range_to"`
	RateAmount           float64             `json:"rate_amount"`
	IsFree               bool                `json:"is_free"`
	Enabled              bool                `json:"enabled"`
	CreatedAt            time.Time           `json:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at"`
	DeliveryMethodType   *DeliveryMethodType `json:"delivery_method_type"`
	Session              *Session            `json:"-"`
}

// DeliveryMethodSession represents the constructor.
type DeliveryMethodSession struct {
	Session
}

// DeliveryMethods creates and returns the constructor.
func (api *API) DeliveryMethods() DeliveryMethodSession {
	var s DeliveryMethodSession
	s.API = api
	s.URI = "delivery_methods"
	s.Options = map[string]string{}
	return s
}

// Page sets the paging option.
func (sess DeliveryMethodSession) Page(page int) DeliveryMethodSession {
	sess.Options["page"] = strconv.Itoa(page)
	return sess
}

// Limit sets the limit option.
func (sess DeliveryMethodSession) Limit(limit int) DeliveryMethodSession {
	sess.Options["limit"] = strconv.Itoa(limit)
	return sess
}

// Find returns a single instance by ID.
func (sess DeliveryMethodSession) Find(id int64) (object *DeliveryMethod, vendenaError *Error) {
	object = &DeliveryMethod{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess DeliveryMethodSession) All() (objects []DeliveryMethod, vendenaError *Error) {
	objects = []DeliveryMethod{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess DeliveryMethodSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess DeliveryMethodSession) New() DeliveryMethod {
	return DeliveryMethod{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *DeliveryMethod) Save() (vendenaError *Error) {
	_, vendenaError = save(&object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *DeliveryMethod) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.ID)
	return
}
