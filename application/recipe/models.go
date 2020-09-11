package recipe

// Recipe - Internal Recipe model
type Recipe struct {
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Link        string `json:"link"`
	Gif         string `json:"gif"`
}

// Response - Internal API response
type Response struct {
	Keywords []string  `json:"keywords"`
	Recipes  []*Recipe `json:"recipes"`
}

// Request - Internal API request
type Request struct {
	Keywords []string `json:"keywords"`
}

// NewRecipe - Create new Recipe model
func NewRecipe(title string, ingredients string, link string, gif string) *Recipe {
	return &Recipe{
		title,
		ingredients,
		link,
		gif,
	}
}

// NewResponse - Create new response model
func NewResponse(keywords []string, recipes []*Recipe) *Response {
	return &Response{
		keywords,
		recipes,
	}
}

// NewRequest - Create new request model
func NewRequest(keywords []string) *Request {
	return &Request{
		keywords,
	}
}
