package secondary

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPConnector Interface for connections
type HTTPConnector interface {
	Do(req *http.Request) (*http.Response, error)
	DoGet(params string) (*http.Response, error)
}

// Connector - Used for HTTP connections
type Connector struct {
	url     string
	headers map[string]string
}

// GetBodyResponse (res)
func GetBodyResponse(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		return nil, err
	}
	return body, nil
}

// NewConnector - Define base URL and Headers for the connection
func NewConnector(url string, headers map[string]string) HTTPConnector {
	return &Connector{
		url:     url,
		headers: headers,
	}
}

// Do - make request
func (c *Connector) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}

// DoGet - make GET request
func (c *Connector) DoGet(params string) (*http.Response, error) {
	return c.getResponse(http.MethodGet, params, nil)
}

func (c *Connector) getResponse(method string, params string, body []byte) (*http.Response, error) {
	url := c.url
	if params != "" {
		url = fmt.Sprintf("%s%s", url, params)
	}
	log.Println("URL", url)
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for i, v := range c.headers {
		r.Header.Set(i, v)
	}
	client := &http.Client{}
	res, err := client.Do(r)

	if res.StatusCode != http.StatusOK {
		return res, err
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}
