package vendena

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const colorGray = "\x1b[30;1m"
const colorRed = "\x1b[31;1m"
const colorGreen = "\x1b[32;1m"
const colorYellow = "\x1b[33;1m"
const colorBlue = "\x1b[34;1m"
const colorMagenta = "\x1b[35;1m"
const colorCyan = "\x1b[36;1m"
const colorWhite = "\x1b[37;1m"
const colorDefault = "\x1b[0m"

// API represents the client configuration.
type API struct {
	URI               string
	BasicAuthUsername string
	BasicAuthPassword string
	ClientID          string
	ClientSecret      string
	StoreID           int64
	ChannelID         int64
	LogURI            bool
	client            *http.Client
}

// SetBasicAuthentication sets the Basic Authentication credentials and removes the API keys.
func (api *API) SetBasicAuthentication(username string, password string) {
	api.BasicAuthUsername = username
	api.BasicAuthPassword = password
	api.ClientID = ""
	api.ClientSecret = ""
}

// SetAPIKeys sets the API keys and removes the Basic Authentication credentials.
func (api *API) SetAPIKeys(clientID string, clientSecret string) {
	api.ClientID = clientID
	api.ClientSecret = clientSecret
	api.BasicAuthUsername = ""
	api.BasicAuthPassword = ""
}

// Session represents the data for the model objects.
type Session struct {
	API     *API
	URI     string
	Options map[string]string
}

// Count is the JSON count response.
type Count struct {
	Count int `json:"count"`
}

// Error represents a custom error object which all operations return.
type Error struct {
	Error        string         `json:"error"`
	Errors       []ErrorDetails `json:"errors"`
	ResponseCode int            `json:"response_code"`
}

// ErrorDetails represents an object containing error details for a VendenaError.
type ErrorDetails struct {
	Resource string `json:"resource"`
	Key      string `json:"key"`
	Message  string `json:"message"`
}

func createError(name string, err error) *Error {
	return &Error{
		Error: name,
		Errors: []ErrorDetails{{
			Resource: "",
			Key:      "",
			Message:  err.Error(),
		}},
	}
}

func parseVendenaError(result *bytes.Buffer, status int) *Error {
	var vendenaError = &Error{}
	if err := json.NewDecoder(result).Decode(vendenaError); err != nil {
		vendenaError = createError("unexpected_error", err)
	}
	vendenaError.ResponseCode = status
	return vendenaError
}

func request(session interface{}, method string, id string, suffix string, body io.Reader) (result *bytes.Buffer, status int, vendenaError *Error) {
	var s = session.(Session)
	var api = s.API
	var options = s.Options
	var endpoint = s.URI

	if api.client == nil {
		api.client = &http.Client{}
	}

	if len(id) > 0 {
		endpoint = fmt.Sprintf("%s/%s", endpoint, id)
	}

	if len(suffix) > 0 {
		endpoint = fmt.Sprintf("%s/%s", endpoint, suffix)
	}

	if len(options) > 0 {
		var values = []string{}
		for k, v := range options {
			values = append(values, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
		endpoint += "?" + strings.Join(values, "&")
	}

	var uri = fmt.Sprintf("%s/%s", api.URI, endpoint)
	if api.LogURI {
		log.Println(uri + " ")
	}

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		vendenaError = createError("network_error", err)
		return
	}

	// Set BasicAuth credentials (if defined)
	if len(api.BasicAuthUsername) > 0 && len(api.BasicAuthPassword) > 0 {
		req.SetBasicAuth(api.BasicAuthUsername, api.BasicAuthPassword)
	}

	// Set API keys (if defined)
	if len(api.ClientID) > 0 && len(api.ClientSecret) > 0 {
		req.SetBasicAuth(api.ClientID, api.ClientSecret)
	}

	// Determine the Store and Channel headers
	if api.StoreID > 0 {
		req.Header.Add("X-Store-ID", strconv.FormatInt(api.StoreID, 10))
	}
	if api.ChannelID > 0 {
		req.Header.Add("X-Channel-ID", strconv.FormatInt(api.ChannelID, 10))
	}

	// Set JSON header
	req.Header.Add("Content-Type", "application/json")

	// Do the actual request
	resp, err := api.client.Do(req)
	// fmt.Printf("resp %v err %v", resp, err)
	if err != nil {
		vendenaError = createError("network_error", err)
		return
	}

	// Return the response status
	status = resp.StatusCode

	// Copy the response stream so we can return the response body
	result = &bytes.Buffer{}
	if _, err = io.Copy(result, resp.Body); err != nil {
		return
	}

	return
}

func findOneByToken(object interface{}, session Session, token string) (status int, vendenaError *Error) {
	result, status, vendenaError := request(session, http.MethodGet, token, "", nil)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&object); err != nil {
		vendenaError = createError("json_decoder_error", err)
	}

	return
}

func findOne(object interface{}, session Session, id int64) (status int, vendenaError *Error) {
	return findOneByToken(object, session, strconv.FormatInt(id, 10))
}

func findAll(objects interface{}, session Session) (status int, vendenaError *Error) {
	result, status, vendenaError := request(session, http.MethodGet, "", "", nil)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&objects); err != nil {
		vendenaError = createError("json_decoder_error", err)
	}

	return
}

func count(session interface{}) (total int, status int, vendenaError *Error) {
	result, status, vendenaError := request(session, http.MethodGet, "", "count", nil)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	var object = &Count{}
	if err := json.NewDecoder(result).Decode(&object); err != nil {
		vendenaError = createError("json_decoder_error", err)
	}

	total = object.Count

	return
}

func saveByToken(object interface{}, session interface{}, token string) (status int, vendenaError *Error) {
	var method string
	var expectedStatus int

	var body = &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(object); err != nil {
		vendenaError = createError("json_encoder_error", err)
		return
	}

	// Determine method based on whether the item is being created or updated
	if token == "" || token == "0" {
		method = http.MethodPost
		expectedStatus = http.StatusCreated
		token = ""
	} else {
		method = http.MethodPut
		expectedStatus = http.StatusOK
	}

	result, status, vendenaError := request(session, method, token, "", body)
	if vendenaError != nil {
		return
	}

	if status != expectedStatus {
		vendenaError = parseVendenaError(result, status)
		return
	}

	if err := json.NewDecoder(result).Decode(&object); err != nil {
		vendenaError = createError("json_decoder_error", err)
		return
	}

	return
}

func save(object interface{}, session interface{}, id int64) (status int, vendenaError *Error) {
	return saveByToken(object, session, strconv.FormatInt(id, 10))
}

func delete(session interface{}, id int64) (status int, vendenaError *Error) {
	result, status, vendenaError := request(session, http.MethodDelete, strconv.FormatInt(id, 10), "", nil)
	if vendenaError != nil {
		return
	}

	if status != http.StatusOK {
		vendenaError = parseVendenaError(result, status)
		return
	}

	return
}
