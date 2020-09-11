package recipe

// Model - Recipe model
type Model struct {
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Link        string `json:"link"`
	Gif         string `json:"gif"`
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
