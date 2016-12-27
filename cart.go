package vendena

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// The Cart model.
type Cart struct {
	ID         int64      `json:"id"`
	ChannelID  int64      `json:"channel_id"`
	Token      string     `json:"token"`
	TotalPrice float64    `json:"total_price"`
	Weight     float64    `json:"weight"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	LineItems  []LineItem `json:"line_items"`
	Session    *Session   `json:"-"`
}

// CartSession represents the constructor.
type CartSession struct {
	Session
}

// Carts creates and returns the constructor.
func (api *API) Carts() CartSession {
	var s CartSession
	s.API = api
	s.URI = "carts"
	s.Options = map[string]string{}
	return s
}

// FindByToken returns a single instance by token.
func (sess CartSession) FindByToken(token string) (object *Cart, err error) {
	_, err = findOneByToken(object, sess.Session, token)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess CartSession) All() (objects []Cart, err error) {
	objects = []Cart{}
	_, err = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess CartSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess CartSession) New() Cart {
	return Cart{
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Cart) Save() (err error) {
	_, err = saveByToken(object, *object.Session, object.Token)
	return
}

// SaveLineItem adds a new line item to the cart.
func (object *Cart) SaveLineItem(lineItem *LineItem) (err error) {
	var body = &bytes.Buffer{}
	err = json.NewEncoder(body).Encode(lineItem)
	if err != nil {
		return
	}

	result, status, err := request(*object.Session, http.MethodPost, object.Token, "items", body)

	if err != nil {
		return
	}

	if status != http.StatusCreated {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	err = json.NewDecoder(result).Decode(lineItem)

	return
}

// RemoveLineItem removes a line item from the cart.
func (object *Cart) RemoveLineItem(id int64) (err error) {
	_, status, err := request(*object.Session, http.MethodDelete, object.Token, "items/"+strconv.FormatInt(id, 10), nil)

	if err != nil {
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	return
}

// Checkout creates a new order from the cart.
func (object *Cart) Checkout() (order Order, err error) {
	var api = object.Session.API
	order = api.Orders().New()

	result, status, err := request(*object.Session, http.MethodPost, object.Token, "checkout", nil)

	if err != nil {
		return
	}

	if status != http.StatusCreated {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	err = json.NewDecoder(result).Decode(&order)

	return
}
