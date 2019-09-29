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

// This function replaces "video" with "library" and strips off a .mp4 string
func convertURLToReadablePath(u string) string {
	filepathSplit := strings.Split(strings.Replace(u, "video", "library", -1), "/")
	filepathSplit = filepathSplit[:len(filepathSplit)-1]
	filepath := strings.Replace(strings.Join(filepathSplit, "/"), "%20", " ", -1)

	return filepath
}

func servePath(w http.ResponseWriter, r *http.Request) {
	// User shouldn't be able to look through the root of the project, that's why it redirects them to /library
	println("URL: ", r.URL.String())
	if r.URL.String() == "/" {
		http.Redirect(w, r, "http://localhost:8080/library/", http.StatusSeeOther)
	} else if strings.Contains(r.URL.String(), "/video") == true {
		t, _ := template.ParseFiles("video.html")
		t.Execute(w, readDir(convertURLToReadablePath(r.URL.String())))
	} else {
		if r.URL.String() != "/favicon.ico" {
			withSpaceChar := strings.Replace(r.URL.String(), "%20", " ", -1)
			// TODO: Check for more file extensions
			if strings.Contains(r.URL.String(), ".mp4") {
				http.ServeFile(w, r, "."+withSpaceChar)
			} else {
				t, _ := template.ParseFiles("library.html")
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
