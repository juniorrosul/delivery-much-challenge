package recipepuppy

// RequestModel - Request model
type RequestModel struct {
	Ingredients []string `json:"ingredients"`
}

// ResponseIndividualModel - Individual recipe response model
type ResponseIndividualModel struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Ingredients string `json:"ingredients"`
	Thumbnail   string `json:"thumbnail"`
}

// ResponseModel - Response model
type ResponseModel struct {
	Recipes []ResponseIndividualModel `json:"results"`
}

// NewRequestModel - New request model
func NewRequestModel(ingredients []string) *RequestModel {
	return &RequestModel{
		ingredients,
	}
}
