package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	ErrInvalidResponse = errors.New("the response is invalid")
)

// Request ...
type CURL struct {
	Request  *http.Request
	Response *http.Response
	Error    error
}

func (c *CURL) SetURLQuery(value url.Values) *CURL {
	if c.Error != nil {
		return c
	}
	c.Request.URL.RawQuery = value.Encode()
	return c
}

// SetHeader ...
func (c *CURL) SetHeader(key, value string) *CURL {
	if c.Error != nil {
		return c
	}
	c.Request.Header.Set(key, value)
	return c
}

// SetHeaderList ...
func (c *CURL) SetHeaderList(vstr ...string) *CURL {
	if c.Error != nil {
		return c
	}
	n := len(vstr)
	for i := 0; i < n; i += 2 {
		c.Request.Header.Set(vstr[i], vstr[i+1])
	}
	return c
}

// Send ...
func (c *CURL) Send() *CURL {
	if c.Error != nil {
		return c
	}
	c.Response, c.Error = http.DefaultClient.Do(c.Request)
	return c
}

// ReadBytes ...
func (c *CURL) ReadBytes() ([]byte, error) {
	if c.Error != nil {
		return nil, c.Error
	}
	if c.Response == nil {
		return nil, ErrInvalidResponse
	}
	data, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ReadJSON ...
func (c *CURL) ReadJSON(x interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(x)
}

// New ...
func New(method, url string, body io.Reader) *CURL {
	req, err := http.NewRequest(method, url, body)

	return &CURL{
		Request: req,
		Error:   err,
	}
}

// Post ...
func Post(url string, body io.Reader) *CURL {
	return New(http.MethodPost, url, body)
}

// PostJSON ...
func PostJSON(url string, x interface{}) *CURL {
	data, err := json.Marshal(x)
	if err != nil {
		return &CURL{Error: err}
	}
	return Post(url, bytes.NewBuffer(data)).SetHeader("Content-Type", "application/json")
}

// Get ...
func Get(url string) *CURL {
	return New(http.MethodGet, url, nil)
}
