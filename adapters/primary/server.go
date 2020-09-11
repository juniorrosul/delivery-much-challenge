package primary

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/juniorrosul/delivery-much-challenge/adapters/primary/middleware"
	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary"
	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary/rpa"
	"github.com/juniorrosul/delivery-much-challenge/application/recipe"
	"github.com/juniorrosul/delivery-much-challenge/application/recipepuppy"
)

func StartServer() {
	log.Println("Starting server...")
	r := mux.NewRouter()
	r.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		keywords := strings.Split(r.URL.Query().Get("i"), ",")
		recipeRequest := recipe.NewRequest(keywords)
		requestModel := recipepuppy.NewRequestModel(recipeRequest.Keywords)
		rpaIntegration := rpa.NewRecipePuppyIntegration(secondary.NewConnector("http://www.recipepuppy.com/api", map[string]string{"Content-Type": "application/json"}))
		_, err := rpaIntegration.GetRecipes(requestModel)
		if err != nil {
			log.Fatal("Error:", err)
		}
	}).Methods("GET")
	r.Use(middleware.EvaluateParameters)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
