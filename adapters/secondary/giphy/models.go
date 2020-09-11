package giphy

import "github.com/juniorrosul/delivery-much-challenge/adapters/secondary"

// GifRequest - Model used for gif request
type GifRequest struct {
	APIKey    string `json:"api"`
	QueryTerm string `json:"q"`
	Limit     int    `json:"limit"`
}

// GifDataResponse - Reponse data of external API
type GifDataResponse struct {
	ID string `json:"id"`
}

// GifResponse - Response of external API
type GifResponse struct {
	Data []GifDataResponse `json:"data"`
}

// Integration - Giphy integration model
type Integration struct {
	conn secondary.HTTPConnector
}

// NewGifRequest -
func NewGifRequest(apiKey string, queryParam string, limit int) *GifRequest {
	return &GifRequest{
		apiKey,
		queryParam,
		limit,
	}
}
