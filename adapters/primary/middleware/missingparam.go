package middleware

import (
	"log"
	"net/http"
	"strings"
)

func EvaluateParameters(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paramString := r.URL.Query().Get("i")
		params := []string{}

		if len(paramString) > 0 {
			params = strings.Split(paramString, ",")
		}

		log.Println("Middleware: EvaluateParameters", r.URL, params)

		if len(params) < 1 || len(params) > 3 {
			log.Println("Middleware: EvaluateParameters: Missing parameters")
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
