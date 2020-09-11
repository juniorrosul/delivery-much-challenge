package recipe

type RecipeModel struct {
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Link        string `json:"link"`
	Gif         string `json:"gif"`
}

type RecipePuppyResponse struct {
	Recipe []RecipeModel `json:"results"`
}

func NewRecipe(title string, ingredients string, link string, gif string) *RecipeModel {
	return &RecipeModel{
		title,
		ingredients,
		link,
		gif,
	}
}
