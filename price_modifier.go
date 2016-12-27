package vendena

// The PriceModifier model.
type PriceModifier struct {
	ID      int      `json:"id"`
	Code    string   `json:"code"`
	Session *Session `json:"-"`
}

// PriceModifierSession represents the constructor.
type PriceModifierSession struct {
	Session
}

// PriceModifiers creates and returns the constructor.
func (api *API) PriceModifiers() PriceModifierSession {
	var s PriceModifierSession
	s.API = api
	s.URI = "price_modifiers"
	s.Options = map[string]string{}
	return s
}

// New creates a new empty object.
func (sess PriceModifierSession) New() PriceModifier {
	return PriceModifier{Session: &sess.Session}
}

// Find returns a single instance by ID.
func (sess PriceModifierSession) Find(id int64) (object *PriceModifier, err error) {
	object = &PriceModifier{}
	_, err = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess PriceModifierSession) All() (objects []PriceModifier, err error) {
	objects = []PriceModifier{}
	_, err = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}
