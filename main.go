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
	// User shouldn't be able to look through the root of the project, that's why it redirects them to /library
	if r.URL.String() == "/" {
		http.Redirect(w, r, "http://localhost:8080/library/", http.StatusSeeOther)
	} else {
		if r.URL.String() != "/favicon.ico" {
			// TODO: Check for more file extensions
			if strings.Contains(r.URL.String(), ".mp4") {
				withSpaceChar := strings.Replace(r.URL.String(), "%20", " ", -1)
				fmt.Println(withSpaceChar)
				http.ServeFile(w, r, "."+withSpaceChar)
			} else {
				t, _ := template.ParseFiles("library.html")
				withSpaceChar := strings.Replace(r.URL.String(), "%20", " ", -1)
				fmt.Println(withSpaceChar)
				http.ServeFile(w, r, "."+withSpaceChar)
				t.Execute(w, readDir(withSpaceChar))
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
