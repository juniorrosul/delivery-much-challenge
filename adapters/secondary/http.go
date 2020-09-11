package secondary

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpConnector interface {
	Do(req *http.Request) (*http.Response, error)
	DoGet(params string) (*http.Response, error)
}

type Connector struct {
	url     string
	headers map[string]string
}

func getBodyResponse(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		fmt.Printf("unexpected status code %d\n body: %s\n", res.StatusCode, string(body))
		return nil, err
	}

	return body, nil
}

func (c *Connector) getResponse(method string, params string, body []byte) (*http.Response, error) {
	url := c.url

	if params != "" {
		url = fmt.Sprintf("%s%s", url, params)
	}

	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	for i, v := range c.headers {
		r.Header.Set(i, v)
	}

	client := &http.Client{}

	res, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewConnector(url string, headers map[string]string) HttpConnector {
	return &Connector{
		url:     url,
		headers: headers,
	}
}

func (c *Connector) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}

func (c *Connector) DoGet(params string) (*http.Response, error) {
	return c.getResponse(http.MethodGet, params, nil)
}
