package giphy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary"
)

// API interface
type API interface {
	searchGif(gr GifRequest)
}

// NewIntegration - Integration initializer
func NewIntegration(httpConnector secondary.HTTPConnector) *Integration {
	return &Integration{
		conn: httpConnector,
	}
}

// GetGif - Get gif endpoint
func (gi *Integration) GetGif(req *GifRequest) (*GifResponse, error) {
	var gr *GifResponse

	res, err := gi.conn.DoGet(fmt.Sprintf("/gifs/search?api_key=%s&q=%s&limit=%d", req.APIKey, url.QueryEscape(req.QueryTerm), req.Limit))
	log.Println(err, res.StatusCode)
	if res != nil && res.StatusCode != http.StatusOK {
		log.Println("Error")
	}

	if err != nil && res == nil {
		return nil, err
	}

	body, err := secondary.GetBodyResponse(res)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		if err := json.Unmarshal(body, &gr); err != nil {
			return nil, err
		}
	}

	return gr, nil
}
