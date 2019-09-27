package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

type pathToServe struct {
	Path string `json:"path"`
}

func readDir(url string) []string {
	dirs := []string{}

	files, err := ioutil.ReadDir("./" + url)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}

func servePath(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/" {
		http.Redirect(w, r, "http://localhost:8080/library/", http.StatusSeeOther)
	} else {
		if r.URL.String() != "/favicon.ico" {
			extension := strings.Split(r.URL.String(), ".")
			lastElement := extension[len(extension)-1]

			if lastElement == "mp4" {
				// TODO: Serve video file
				fmt.Println("MP4", r.URL.String())
				http.ServeFile(w, r, "."+r.URL.String())
			} else {
				t, _ := template.ParseFiles("library.html")
				t.Execute(w, readDir(r.URL.String()))
			}
		}
	}
}

func main() {
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.PathPrefix("/").HandlerFunc(servePath).Methods("GET")

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r))
}
