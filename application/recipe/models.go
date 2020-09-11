package recipe

// Model - Recipe model
type Model struct {
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Link        string `json:"link"`
	Gif         string `json:"gif"`
}

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
	Recipe []ResponseIndividualModel `json:"results"`
}

// NewRecipe - Create new Recipe model
func NewRecipe(title string, ingredients string, link string, gif string) *Model {
	return &Model{
		title,
		ingredients,
		link,
		gif,
	}
}
