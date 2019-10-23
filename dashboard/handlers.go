package main

import (
	"net/http"
	"strings"
	"text/template"
)

func ServePath(w http.ResponseWriter, r *http.Request) {
	// User shouldn't be able to look through the root of the project, that's why it redirects them to /library
	println("URL: ", r.URL.String())
	if r.URL.String() == "/" {
		http.Redirect(w, r, "http://localhost:8080/library/", http.StatusSeeOther)
	} else if strings.Contains(r.URL.String(), "/video") == true {
		t, _ := template.ParseFiles("video.html")
		t.Execute(w, ReadDir(ConvertURLToReadablePath(r.URL.String())))
	} else {
		if r.URL.String() != "/favicon.ico" {
			withSpaceChar := strings.Replace(r.URL.String(), "%20", " ", -1)
			// TODO: Check for more file extensions
			if strings.Contains(r.URL.String(), ".mp4") {
				http.ServeFile(w, r, "."+withSpaceChar)
			} else {
				t, _ := template.ParseFiles("library.html")
				t.Execute(w, ReadDir(withSpaceChar))
			}
		}
	}
}
