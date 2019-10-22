package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type SuccessfulResponse struct {
	Email string `json:email`
	Token string `json:jwt`
}

type FailedResponse struct {
	Error string `json:error`
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	passwd := r.FormValue("password")
	passwd_c := r.FormValue("password_confirm")

	isEmailValid, err := IsEmailValid(email)
	if err != nil {
		response := FailedResponse{Error: err.Error()}
		ReturnError(response, w)
	}

	isPasswordValid, err := IsPasswordValid(passwd)
	if err != nil {
		response := FailedResponse{Error: err.Error()}
		ReturnError(response, w)
	}

	arePasswordsTheSame, err := ArePasswordsTheSame(passwd, passwd_c)
	if err != nil {
		response := FailedResponse{Error: err.Error()}
		ReturnError(response, w)
	}

	if arePasswordsTheSame == true {
		if isEmailValid == true && isPasswordValid == true {
			b := DoesUserExist(email)

			if b == true {
				response := FailedResponse{Error: "This user already exists"}

				ReturnError(response, w)
			} else {
				hashedPasswd, err := HashPassword(passwd)
				if err != nil {
					log.Fatal(err)
				}

				newUser := User{email, hashedPasswd}

				InsertUser(newUser)

				jwt := GenerateToken(email)

				response := SuccessfulResponse{Email: email, Token: jwt}

				json.NewEncoder(w).Encode(response)
			}
		}
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
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
