package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Started register server")

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/register", HandleRegister).Methods("POST")
	r.HandleFunc("/", HandleIndex).Methods("GET")
	r.PathPrefix("/js").HandlerFunc(ServeJS).Methods("GET")
	r.PathPrefix("/css").HandlerFunc(ServeCSS).Methods("GET")
	r.PathPrefix("/fonts").HandlerFunc(ServeFonts).Methods("GET")
	r.PathPrefix("/vendor").HandlerFunc(ServeVendor).Methods("GET")

	http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(r))

	fmt.Println("Connected to Mongo")
}
