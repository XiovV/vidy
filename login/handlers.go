package main

import (
	"encoding/json"
	"net/http"
)

type SuccessfulResponse struct {
	Email string `json:"email"`
	Token string `json:jwt`
}

type FailedResponse struct {
	Error string
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := r.FormValue("email")
	passwd := r.FormValue("password")

	b := DoesUserExist(email)

	if b == true {
		areCredsGood, err := CheckCredentials(email, passwd)
		if err != nil {
			response := FailedResponse{Error: "Email or Password is incorrect"}
			json.NewEncoder(w).Encode(response)
		}

		if areCredsGood == true {
			jwt := GenerateToken(email)

			response := SuccessfulResponse{Email: email, Token: jwt}
			json.NewEncoder(w).Encode(response)
		} else {
			response := FailedResponse{Error: "Email or Password is incorrect"}
			json.NewEncoder(w).Encode(response)
		}
	} else {
		response := FailedResponse{Error: "This user doesn't exist"}
		json.NewEncoder(w).Encode(response)
	}

}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "static/index.html")
}

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "static"+r.URL.String())
}

func ServeVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "static"+r.URL.String())
}

func ServeJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "static"+r.URL.String())
}

func ServeFonts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "static"+r.URL.String())
}
