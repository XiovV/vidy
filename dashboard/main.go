package main

import (
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.PathPrefix("/library").HandlerFunc(ServePath).Methods("GET")
	r.PathPrefix("/video").HandlerFunc(ServePath).Methods("GET")

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r))
}
