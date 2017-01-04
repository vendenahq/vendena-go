package vendena

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	vendena "kvika.io/vendena-go"

	"github.com/stretchr/testify/assert"
)

var (
	api = API{
		URI:       "http://localhost:8080",
		StoreID:   int64(0),
		ChannelID: int64(0),
		LogURI:    false,
	}
	deliveryMethod = api.DeliveryMethods().New()
	paymentMethod  = api.PaymentMethods().New()
	productOption  = api.ProductOptions().New()
	image          = api.Images().New()
	product        = api.Products().New()
	cart           = api.Carts().New()
	lineItem       = api.LineItems().New()
	order          = api.Orders().New()
	allowCreate    = true
)

func init() {
	if os.Getenv("VENDENA_API_HOST") != "" {
		api.URI = os.Getenv("VENDENA_API_HOST")
	}

	if os.Getenv("VENDENA_API_USERNAME") != "" && os.Getenv("VENDENA_API_PASSWORD") != "" {
		api.SetBasicAuthentication(os.Getenv("VENDENA_API_USERNAME"), os.Getenv("VENDENA_API_PASSWORD"))
	}

	if os.Getenv("VENDENA_API_CLIENT_ID") != "" && os.Getenv("VENDENA_API_CLIENT_SECRET") != "" {
		api.SetAPIKeys(os.Getenv("VENDENA_API_CLIENT_ID"), os.Getenv("VENDENA_API_CLIENT_SECRET"))
	}

	if os.Getenv("VENDENA_API_STORE_ID") != "" {
		if id, err := strconv.ParseInt(os.Getenv("VENDENA_API_STORE_ID"), 10, 64); err == nil {
			api.StoreID = id
		}
	}

	if os.Getenv("VENDENA_API_CHANNEL_ID") != "" {
		if id, err := strconv.ParseInt(os.Getenv("VENDENA_API_CHANNEL_ID"), 10, 64); err == nil {
			api.ChannelID = id
		}
	}
}

func Describe(title string) {
	// log.Println("---")
	// log.Println(colorWhite + title + colorDefault + " ")
	fmt.Print(title + " ")
}

func Done(t *testing.T) {
	if t.Failed() {
		// log.Println(colorRed + "X Failed" + colorDefault)
		fmt.Println(colorRed + "X" + colorDefault)
	} else {
		// log.Println(colorGreen + "√ Passed" + colorDefault)
		fmt.Println(colorGreen + "√" + colorDefault)
	}
}

//
// Payment Gateway
//

func TestGetPaymentGateways(t *testing.T) {
	Describe("Get Payment Gateways")

	var gateway, err = api.PaymentGateways().All()
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, gateway, "should have been found")
	assert.NotEqual(t, len(gateway), 0, "should not be equal")

	Done(t)
}

func TestGetPaymentGateway(t *testing.T) {
	Describe("Get Payment Gateway")

	var gateway, err = api.PaymentGateways().Find(1)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, gateway, "should have been found")
	assert.NotEmpty(t, gateway.Code, "Code should not be empty")

	Done(t)
}

func TestCountPaymentGateway(t *testing.T) {
	Describe("Count Payment Gateways")

	var count, err = api.PaymentGateways().Count()
	assert.Nil(t, err, "should be nil")
	assert.NotEqual(t, count, 0, "should not be equal")

	Done(t)
}

//
// Stores
//

func TestGetStore(t *testing.T) {
	Describe("Get Store")

	var store, err = api.Stores().Find(1)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, store, "should have been found")
	assert.NotEmpty(t, store.Title, "Title should not be empty")

	Done(t)
}

//
// Channels
//

// func TestCreateChannel(t *testing.T) {
// 	if allowCreate {
// 		Describe("Create Channel")
//
// 		var channel = api.Channels().New()
// 		channel.Title = "Created by SDK"
// 		channel.DefaultCurrencyID = 2
// 		channel.DefaultLocaleID = 2
// 		var err = channel.Save()
// 		assert.Nil(t, err, "should be nil")
// 		assert.NotEqual(t, channel.ID, int64(0), "ID should not be 0")
// 		assert.NotEmpty(t, channel.Title, "Title should not be empty")
//
// 		Done(t)
// 	}
// }

//
// Delivery Methods
//

func TestCreateDeliveryMethod(t *testing.T) {
	if allowCreate {
		Describe("Create Delivery Method")

		deliveryMethod.DeliveryMethodTypeID = int64(1)
		deliveryMethod.Title = "Price Based Shipping"
		deliveryMethod.RangeFrom = 10.00
		deliveryMethod.RangeTo = 50.00
		deliveryMethod.RateAmount = 5.95
		deliveryMethod.IsFree = false

		var err = deliveryMethod.Save()
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, deliveryMethod.ID, int64(0), "ID should not be 0")
		assert.NotEmpty(t, deliveryMethod.Title, "Title should not be empty")

		Done(t)
	}
}

//
// Payment Methods
//

func TestCreatePaymentMethod(t *testing.T) {
	if allowCreate {
		Describe("Create Payment Method")

		paymentMethod.PaymentMethodTypeID = int64(1)
		paymentMethod.PaymentGatewayID = int64(1)
		paymentMethod.Title = "Credit Card"
		paymentMethod.Details = ""
		paymentMethod.Instructions = ""
		paymentMethod.Fee = 0.5

		var err = paymentMethod.Save()
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, paymentMethod.ID, int64(0), "ID should not be 0")
		assert.NotEmpty(t, paymentMethod.Title, "Title should not be empty")

		Done(t)
	}
}

//
// Taxons
//

func TestGetTaxon(t *testing.T) {
	Describe("Get Taxon")

	var taxon, err = api.Taxons().Find(1)
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, taxon.Code, "category", "should be equal")

	Done(t)
}

func TestGetTaxons(t *testing.T) {
	Describe("Get Taxons")

	var taxons, err = api.Taxons().All()
	assert.Nil(t, err, "should be nil")
	assert.NotEqual(t, len(taxons), 0, "should not be equal")

	Done(t)
}

// func TestCountTaxons(t *testing.T) {
// 	Describe("Count Taxons")
//
// 	var count, err = api.Taxons().Count()
// 	assert.Nil(t, err, "should be nil")
// 	assert.Equal(t, count, 4, "should be equal")
//
// 	Done(t)
// }

//
// Product Options
//

func TestCreateProductOption(t *testing.T) {
	if allowCreate {
		Describe("Create Product Option")

		productOption.Code = "tshirt-size"
		productOption.Title = "T-shirt Size"

		var pov1 = api.ProductOptionValues().New()
		pov1.Title = "Small"
		productOption.AddValue(pov1)

		var pov2 = api.ProductOptionValues().New()
		pov2.Title = "Medium"
		productOption.AddValue(pov2)

		var err = productOption.Save()
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, productOption.ID, int64(0), "ID should not be 0")
		assert.NotEmpty(t, productOption.Title, "Title should not be empty")
		assert.Equal(t, len(productOption.Values), 2, "should be equal")
		assert.Equal(t, productOption.Values[0].Title, pov1.Title, "should be equal")
		assert.Equal(t, productOption.Values[1].Title, pov2.Title, "should be equal")

		Done(t)
	}
}

func TestUpdateProductOption(t *testing.T) {
	if allowCreate {
		Describe("Update Product Option")

		productOption.Title = "Size"

		var err = productOption.Save()
		assert.Nil(t, err, "should be nil")
		assert.Equal(t, productOption.Title, "Size", "should be equal")

		Done(t)
	}
}

//
// Images
//

func TestCreateImage(t *testing.T) {
	if allowCreate {
		Describe("Create Image")

		image.URL = "https://upload.wikimedia.org/wikipedia/commons/7/79/Frankie_Say_War%21_Hide_Yourself%22_t-shirt.jpg"

		var err = image.Save()
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, image.ID, int64(0), "ID should not be 0")
		assert.NotEmpty(t, image.URL, "URL should not be empty")

		Done(t)
	}
}

// //
// // Price Modifiers
// //
//
// func TestGetPriceModifier(t *testing.T) {
// 	Describe("Get Price Modifier")
//
// 	var modifier, err = api.PriceModifiers().Find(1)
//
// 	assert.Nil(t, err, "should be nil")
// 	assert.NotEqual(t, modifier.ID, int64(0), "ID should not be 0")
// 	assert.NotEmpty(t, modifier.Code, "Code should not be empty")
//
// 	Done(t)
// }

//
// Products
//

func TestCreateProduct(t *testing.T) {
	if allowCreate {
		Describe("Create Product")

		product.Title = "Created by SDK"
		product.Price = 20.0
		product.Weight = 0.5
		product.StockLevel = 10
		product.AddImage(image)

		var variant1 = api.Variants().New()
		variant1.ProductOptionValueIDs = []int64{productOption.Values[0].ID, productOption.Values[1].ID}
		variant1.PriceModifierID = int64(1)
		variant1.Price = 10.0
		variant1.StockLevel = 10
		variant1.AddImage(image)
		product.AddVariant(variant1)

		var variant2 = api.Variants().New()
		variant2.Title = "No option values"
		variant2.PriceModifierID = int64(2)
		variant2.Price = 10.0
		variant2.StockLevel = 10
		product.AddVariant(variant2)

		var err = product.Save()
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, product.ID, int64(0), "ID should not be 0")
		assert.NotEmpty(t, product.Title, "Title should not be empty")

		Done(t)
	}
}

func TestGetProduct(t *testing.T) {
	Describe("Get Product")

	var product, err = api.Products().Find(2)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, product, "should have been found")
	assert.NotEmpty(t, product.Title, "Title should not be empty")

	Done(t)
}

func TestGetProducts(t *testing.T) {
	Describe("Get Products")

	var products, err = api.Products().Page(1).Limit(2).Taxons(2, 4).All()
	assert.Nil(t, err, "should be nil")
	assert.NotEqual(t, len(products), 0, "should not be equal")

	Done(t)
}

func TestCountProducts(t *testing.T) {
	Describe("Count Products")

	var count, err = api.Products().Count()
	assert.Nil(t, err, "should be nil")
	assert.NotEqual(t, count, 0, "should not be equal")

	Done(t)
}

//
// Carts
//

func TestCreateCart(t *testing.T) {
	if allowCreate {
		Describe("Create Cart")

		var err = cart.Save()
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, cart.ID, int64(0), "ID should not be 0")
		assert.NotEmpty(t, cart.Token, "Token should not be empty")

		Done(t)
	}
}

func TestAddItemsToCart(t *testing.T) {
	if allowCreate {
		Describe("Add Items To Cart")

		var err *vendena.Error

		lineItem.ProductID = product.ID
		lineItem.ProductOptionValueIDs = []int64{product.Variants[0].ProductOptionValueIDs[0], product.Variants[0].ProductOptionValueIDs[1]}
		lineItem.Quantity = 1

		var lineItem1 = api.LineItems().New()
		lineItem1.ProductID = product.ID
		lineItem1.ProductOptionValueIDs = []int64{product.Variants[0].ProductOptionValueIDs[0], product.Variants[0].ProductOptionValueIDs[1]}
		lineItem1.Quantity = 1

		var lineItem2 = api.LineItems().New()
		lineItem2.ProductID = product.ID
		lineItem2.VariantID = product.Variants[1].ID
		lineItem2.Quantity = 2

		err = cart.SaveLineItem(&lineItem)
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, lineItem.ID, int64(0), "ID should not be 0")

		err = cart.SaveLineItem(&lineItem1)
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, lineItem1.ID, int64(0), "ID should not be 0")

		err = cart.SaveLineItem(&lineItem2)
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, lineItem2.ID, int64(0), "ID should not be 0")

		Done(t)
	}
}

func TestRemoveItemFromCart(t *testing.T) {
	if allowCreate {
		Describe("Remove Item From Cart")

		var err = cart.RemoveLineItem(lineItem.ID)
		assert.Nil(t, err, "should be nil")

		Done(t)
	}
}

//
// Orders
//

func TestCheckout(t *testing.T) {
	if allowCreate {
		Describe("Checkout")

		var err error

		order, err = cart.Checkout()
		assert.Nil(t, err, "should be nil")
		assert.NotEqual(t, order.ID, int64(0), "ID should not be 0")

		Done(t)
	}
}

func TestUpdateOrder(t *testing.T) {
	if allowCreate {
		Describe("Update Order")

		var customer = api.Customers().New()
		customer.FirstName = "John"
		customer.LastName = "Doe"
		customer.Email = "john.doe@example.com"

		var address = api.Addresses().New()
		address.FirstName = "John"
		address.LastName = "Doe"
		address.Email = "john.doe@example.com"
		address.Address1 = "123 Street Address"
		address.Address2 = "Apt. 101"
		address.City = "City"
		address.Province = "VA"
		address.Country = "US"
		address.Postcode = "123455"
		address.Phone = "+1 202 555 0162"

		// order.StatusID = int64(1)
		order.DeliveryMethodID = int64(1)
		order.PaymentMethodID = paymentMethod.ID
		order.Customer = &customer
		order.ShippingAddress = &address
		order.UseBillingAddress = false

		var err = order.Save()
		assert.Nil(t, err, "should be nil")
		assert.Equal(t, order.Status.Code, "checkout", "should be equal")
		assert.Equal(t, order.Customer.FirstName, "John", "should be equal")

		Done(t)
	}
}

func TestMarkOrderOpen(t *testing.T) {
	if allowCreate {
		Describe("Mark Order Open")

		var statusOpen = int64(3)
		var statusPaid = int64(4)
		var statusPendingFulfillment = int64(1)

		order.Customer.FirstName = "Jonathan"
		order.StatusID = statusOpen
		order.PaymentStatusID = statusPaid
		order.FulfillmentStatusID = statusPendingFulfillment

		var err = order.Save()
		assert.Nil(t, err, "should be nil")
		assert.Equal(t, order.Customer.FirstName, "Jonathan", "should be equal")
		assert.Equal(t, order.Status.Code, "open", "should be equal")
		assert.Equal(t, order.PaymentStatus.Code, "paid", "should be equal")

		Done(t)
	}
}

func TestGetOrderFormData(t *testing.T) {
	if allowCreate {
		Describe("Get Order Form Data")

		var formData, err = order.FormData()
		assert.Nil(t, err, "should be nil")
		assert.NotEmpty(t, formData.URL, "should not be empty")
		assert.NotEmpty(t, formData.Elements, "should not be empty")

		Done(t)
	}
}

func TestValidatePaymentNotification(t *testing.T) {
	if allowCreate {
		Describe("Validate Payment Notification")

		var result, err = order.ValidateNotification("test=true")
		assert.Nil(t, err, "should be nil")
		assert.Equal(t, result.Valid, false, "should be false")

		Done(t)
	}
}

//
// Clean up
//

func TestDeleteDeliveryMethod(t *testing.T) {
	if allowCreate {
		Describe("Delete Delivery Method")

		var err = deliveryMethod.Delete()
		assert.Nil(t, err, "should be nil")

		Done(t)
	}
}

func TestDeletePaymentMethod(t *testing.T) {
	if allowCreate {
		Describe("Delete Payment Method")

		var err = paymentMethod.Delete()
		assert.Nil(t, err, "should be nil")

		Done(t)
	}
}

func TestDeleteProductOption(t *testing.T) {
	if allowCreate {
		Describe("Delete Product Option")

		var err = productOption.Delete()
		assert.Nil(t, err, "should be nil")

		Done(t)
	}
}

func TestDeleteImage(t *testing.T) {
	if allowCreate {
		Describe("Delete Image")

		var err = image.Delete()
		assert.Nil(t, err, "should be nil")

		Done(t)
	}
}

func TestDeleteProduct(t *testing.T) {
	if allowCreate {
		Describe("Delete Product")

		var err = product.Delete()
		assert.Nil(t, err, "should be nil")

		Done(t)
	}
}
