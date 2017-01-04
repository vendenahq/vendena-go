package vendena

import (
	"strconv"
	"strings"
	"time"
)

// The Product model.
type Product struct {
	ID                   int64     `json:"id"`
	Title                string    `json:"title"`
	Body                 string    `json:"body"`
	Excerpt              string    `json:"excerpt"`
	Slug                 string    `json:"slug"`
	Price                float64   `json:"price"`
	Weight               float64   `json:"weight"`
	SKU                  string    `json:"sku"`
	StockLevel           int       `json:"stock_level"`
	StockLevelLowWarning int       `json:"stock_level_low_warning"`
	Enabled              bool      `json:"enabled"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Images               []Image   `json:"images"`
	Variants             []Variant `json:"variants"`
	Taxons               []Taxon   `json:"taxons"`
	TaxonIDs             []int64   `json:"taxon_ids"`
	Session              *Session  `json:"-"`
}

// ProductSession represents the constructor.
type ProductSession struct {
	Session
}

// Products creates and returns the constructor.
func (api *API) Products() ProductSession {
	var s ProductSession
	s.API = api
	s.URI = "products"
	s.Options = map[string]string{}
	return s
}

// Page sets the paging option.
func (sess ProductSession) Page(page int) ProductSession {
	sess.Options["page"] = strconv.Itoa(page)
	return sess
}

// Limit sets the limit option.
func (sess ProductSession) Limit(limit int) ProductSession {
	sess.Options["limit"] = strconv.Itoa(limit)
	return sess
}

// Taxons sets the taxon_ids option.
func (sess ProductSession) Taxons(values ...int64) ProductSession {
	var valuesText = []string{}
	for i := range values {
		var text = strconv.FormatInt(values[i], 10)
		valuesText = append(valuesText, text)
	}
	sess.Options["taxon_ids"] = strings.Join(valuesText, ",")
	return sess
}

// Find returns a single instance by ID.
func (sess ProductSession) Find(id int64) (object *Product, vendenaError *Error) {
	object = &Product{}
	_, vendenaError = findOne(object, sess.Session, id)
	object.Session = &sess.Session
	return
}

// All returns all instances.
func (sess ProductSession) All() (objects []Product, vendenaError *Error) {
	objects = []Product{}
	_, vendenaError = findAll(&objects, sess.Session)
	for i := range objects {
		objects[i].Session = &sess.Session
	}
	return
}

// Count returns the number of instances.
func (sess ProductSession) Count() (total int, vendenaError *Error) {
	total, _, vendenaError = count(sess.Session)
	return
}

// New creates a new empty object.
func (sess ProductSession) New() Product {
	return Product{
		Enabled: true,
		Session: &sess.Session,
	}
}

// Save creates or updates an object.
func (object *Product) Save() (vendenaError *Error) {
	_, vendenaError = save(object, *object.Session, object.ID)
	return
}

// AddImage adds a new Image.
func (object *Product) AddImage(image Image) {
	object.Images = append(object.Images, image)
}

// AddVariant adds a new Variant.
func (object *Product) AddVariant(variant Variant) {
	object.Variants = append(object.Variants, variant)
}

// Delete deletes an object.
func (object *Product) Delete() (vendenaError *Error) {
	_, vendenaError = delete(*object.Session, object.ID)
	return
}
