package primary

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/juniorrosul/delivery-much-challenge/adapters/primary/middleware"
	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary"
	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary/giphy"
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
		giphyIntegration := giphy.NewIntegration(secondary.NewConnector("http://api.giphy.com/v1", map[string]string{"Content-Type": "application/json"}))

		recipes, err := rpaIntegration.GetRecipes(requestModel)
		if err != nil {
			log.Fatal("Error:", err)
		}

		var response recipe.Response
		response.Keywords = keywords

		for i := 0; i < len(recipes.Recipes); i++ {
			giphyReq := giphy.NewGifRequest(os.Getenv("GIPHY_API_KEY"), recipes.Recipes[i].Title, 1)
			giphyResponse, err := giphyIntegration.GetGif(giphyReq)
			if err != nil {
				log.Fatal("Error: ", err)
			}

			ingredients := strings.Split(recipes.Recipes[i].Ingredients, ", ")

			sort.Strings(ingredients)

			newRecipe := recipe.NewRecipe(
				recipes.Recipes[i].Title,
				ingredients,
				recipes.Recipes[i].Href,
				fmt.Sprintf("https://media.giphy.com/media/%s/giphy.gif", giphyResponse.Data[0].ID),
			)
			response.Recipes = append(response.Recipes, newRecipe)
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("GET")
	r.Use(middleware.EvaluateParameters)

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatal(err)
	}
}
