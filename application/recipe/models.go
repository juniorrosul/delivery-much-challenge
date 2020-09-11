package recipe

type RecipeModel struct {
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Link        string `json:"link"`
	Gif         string `json:"gif"`
}

type RecipePuppyModel struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Ingredients string `json:"ingredients"`
	Thumbnail   string `json:"thumbnail"`
}

type RecipePuppyResponseModel struct {
	Recipe []RecipePuppyModel `json:"results"`
}

func NewRecipe(title string, ingredients string, link string, gif string) *RecipeModel {
	return &RecipeModel{
		title,
		ingredients,
		link,
		gif,
	}
}
