package rpa

import "net/http"

// ConnectorMock - Mock for connector
type ConnectorMock struct {
	ResponseStatus int
	ResponseBody   string
}

// Do - Mock request
func (c *ConnectorMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

// DoGet - Mock request
func (c *ConnectorMock) DoGet(params string) (*http.Response, error) {
	return &http.Response{}, nil
}
