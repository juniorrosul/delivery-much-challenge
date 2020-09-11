package secondary

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

type mock struct{}

func (m *mock) Read(p []byte) (n int, err error) {
	return 0, errors.New("testing error")
}

func (m *mock) Close() error {
	return errors.New("testing error")
}

func TestGetBodyResponse(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		response := &http.Response{}
		response.StatusCode = http.StatusOK
		response.Body = ioutil.NopCloser(bytes.NewBufferString("{\"response\":\"body\"}"))

		_, e := GetBodyResponse(response)
		if e != nil {
			t.Error("expected error")
		}
	})

	t.Run("UNEXPECTED BODY", func(t *testing.T) {
		res := &http.Response{}
		res.Body = &mock{}
		_, e := GetBodyResponse(res)
		if e == nil {
			t.Error("An error was  expected")
		}
	})
}

func TestNewConnector(t *testing.T) {
	t.Run("Do", func(t *testing.T) {
		n := NewConnector("https://postman-echo.com", map[string]string{
			"Origin": "localhost",
		})
		req, _ := http.NewRequest("GET", "https://postman-echo.com", nil)
		res, err := n.Do(req)
		if err != nil {
			t.Error("An error was not expected")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("An %v was expected and return %v", http.StatusOK, res.StatusCode)
		}
	})

	t.Run("DoGet", func(t *testing.T) {
		n := NewConnector("https://postman-echo.com", map[string]string{
			"Origin": "localhost",
		})
		res, err := n.DoGet("/get")
		if err != nil {
			t.Error("An error was not expected")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("An %v was expected and return %v", http.StatusOK, res.StatusCode)
		}
	})

}

func TestCloseConnection(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://postman-echo.com/get", ioutil.NopCloser(bytes.NewBufferString("{\"response\":\"body\"}")))
	r.Body = &mock{}
	CloseConnection(r)
}
