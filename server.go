package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Recipe model
type Recipe struct {
	Title       string `json:"name"`
	Ingredients string `json:"ingredients"`
	Link        string `json:"link"`
	Gif         string `json:"gif"`
}

type recipeHandlers struct {
	sync.Mutex
	store map[string]Recipe
}

func (h *recipeHandlers) get(w http.ResponseWriter, r *http.Request) {
	recipes := make([]Recipe, len(h.store))

	h.Lock()
	i := 0
	for _, recipe := range h.store {
		recipes[i] = recipe
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(recipes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newRecipeHandlers() *recipeHandlers {
	return &recipeHandlers{
		store: map[string]Recipe{
			"Recipe 1": {
				Title:       "name",
				Ingredients: "ingredients",
				Link:        "link",
				Gif:         "gif",
			},
		},
	}
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	recipeHandler := newRecipeHandlers()
	http.HandleFunc("/recipes", recipeHandler.get)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
