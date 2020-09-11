package primary

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juniorrosul/delivery-much-challenge/adapters/primary/middleware"
)

func StartServer() {
	log.Println("Starting server...")

	r := mux.NewRouter()

	r.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
	}).Methods("GET")

	r.Use(middleware.EvaluateParameters)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
