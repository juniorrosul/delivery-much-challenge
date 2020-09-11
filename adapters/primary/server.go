package primary

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/juniorrosul/delivery-much-challenge/adapters/primary/middleware"
	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary"
	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary/rpa"
	"github.com/juniorrosul/delivery-much-challenge/application/recipe"
	"github.com/juniorrosul/delivery-much-challenge/application/recipepuppy"
)

// StartServer - Start server function
func StartServer() {
	log.Println("Starting server...")
	r := mux.NewRouter()
	r.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		keywords := strings.Split(r.URL.Query().Get("i"), ",")
		recipeRequest := recipe.NewRequest(keywords)
		requestModel := recipepuppy.NewRequestModel(recipeRequest.Keywords)
		rpaIntegration := rpa.NewRecipePuppyIntegration(secondary.NewConnector("http://www.recipepuppy.com/api", map[string]string{"Content-Type": "application/json"}))
		recipes, err := rpaIntegration.GetRecipes(requestModel)
		if err != nil {
			log.Fatal("Error:", err)
		}

		var response recipe.Response
		response.Keywords = keywords

		for i := 0; i < len(recipes.Recipes); i++ {
			ingredients := strings.Split(recipes.Recipes[i].Ingredients, ", ")

			sort.Strings(ingredients)

			newRecipe := recipe.NewRecipe(
				recipes.Recipes[i].Title,
				ingredients,
				recipes.Recipes[i].Href,
				recipes.Recipes[i].Thumbnail,
			)
			response.Recipes = append(response.Recipes, newRecipe)
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("GET")
	r.Use(middleware.EvaluateParameters)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
