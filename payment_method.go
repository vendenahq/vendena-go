package vendena

import "time"

// The PaymentMethod model.
type PaymentMethod struct {
	ID                  int64              `json:"id"`
	PaymentMethodTypeID int64              `json:"payment_method_type_id"`
	PaymentGatewayID    int64              `json:"payment_gateway_id"`
	Credentials1        string             `json:"credentials_1"`
	Credentials2        string             `json:"credentials_2"`
	Credentials3        string             `json:"credentials_3"`
	Title               string             `json:"title"`
	Details             string             `json:"details"`
	Instructions        string             `json:"instructions"`
	Fee                 float64            `json:"float"`
	TestMode            bool               `json:"test_mode"`
	Enabled             bool               `json:"enabled"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	PaymentMethodType   *PaymentMethodType `json:"payment_method_type"`
	PaymentGateway      *PaymentGateway    `json:"payment_gateway"`
	Session             *Session           `json:"-"`
}

// PaymentMethodSession represents the constructor.
type PaymentMethodSession struct {
	Session
}

// PaymentMethods creates and returns the constructor.
func (api *API) PaymentMethods() PaymentMethodSession {
	var s PaymentMethodSession
	s.API = api
	s.URI = "payment_methods"
	s.Options = map[string]string{}
	return s
}

// Find returns a single instance by ID.
func (sess PaymentMethodSession) Find(id int64) (object *PaymentMethod, err error) {
	object = &PaymentMethod{}
	_, err = findOne(object, sess.Session, id)
	return
}

// All returns all instances.
func (sess PaymentMethodSession) All() (objects []PaymentMethod, err error) {
	objects = []PaymentMethod{}
	_, err = findAll(&objects, sess.Session)
	return
}

// Count returns the number of instances.
func (sess PaymentMethodSession) Count() (total int, err error) {
	total, _, err = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess PaymentMethodSession) New() PaymentMethod {
	return PaymentMethod{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *PaymentMethod) Save() (err error) {
	_, err = save(object, *object.Session, object.ID)
	return
}

// Delete deletes an object.
func (object *PaymentMethod) Delete() (err error) {
	_, err = delete(*object.Session, object.ID)
	return
}
