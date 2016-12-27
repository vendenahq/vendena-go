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
	URI       string
	Token     string
	Secret    string
	ChannelID int64
	LogURI    bool
	client    *http.Client
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

type errorResponse struct {
	Errors map[string]interface{} `json:"errors"`
}

func request(session interface{}, method string, id string, suffix string, body io.Reader) (result *bytes.Buffer, status int, err error) {
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
		return
	}

	req.SetBasicAuth(api.Token, api.Secret)
	req.Header.Add("X-Channel-ID", strconv.FormatInt(api.ChannelID, 10))
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	// fmt.Printf("resp %v err %v", resp, err)
	if err != nil {
		return
	}

	status = resp.StatusCode

	result = &bytes.Buffer{}
	if _, err = io.Copy(result, resp.Body); err != nil {
		return
	}

	return
}

func findOneByToken(object interface{}, session Session, token string) (status int, err error) {
	result, status, err := request(session, http.MethodGet, token, "", nil)

	if err != nil {
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	err = json.NewDecoder(result).Decode(&object)

	return
}

func findOne(object interface{}, session Session, id int64) (status int, err error) {
	return findOneByToken(object, session, strconv.FormatInt(id, 10))
}

func findAll(objects interface{}, session Session) (status int, err error) {
	result, status, err := request(session, http.MethodGet, "", "", nil)

	if err != nil {
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	err = json.NewDecoder(result).Decode(&objects)

	return
}

func count(session interface{}) (total int, status int, err error) {
	result, status, err := request(session, http.MethodGet, "", "count", nil)

	if err != nil {
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	var object = &Count{}
	err = json.NewDecoder(result).Decode(&object)

	if err != nil {
		return
	}

	total = object.Count

	return
}

func saveByToken(object interface{}, session interface{}, token string) (status int, err error) {
	var method string
	var expectedStatus int

	var body = &bytes.Buffer{}
	err = json.NewEncoder(body).Encode(object)
	if err != nil {
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

	result, status, err := request(session, method, token, "", body)

	if err != nil {
		return
	}

	if status != expectedStatus {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	err = json.NewDecoder(result).Decode(&object)

	return
}

func save(object interface{}, session interface{}, id int64) (status int, err error) {
	return saveByToken(object, session, strconv.FormatInt(id, 10))
}

func delete(session interface{}, id int64) (status int, err error) {
	_, status, err = request(session, http.MethodDelete, strconv.FormatInt(id, 10), "", nil)

	if err != nil {
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("Status returned: %d", status)
		return
	}

	return
}
