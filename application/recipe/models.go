package recipe

type Model struct {
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Link        string `json:"link"`
	Gif         string `json:"gif"`
}
