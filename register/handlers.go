package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
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

	isEmailValid, err := IsEmailValid(email)
	if err != nil {
		response := FailedResponse{Error: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	isPasswordValid, err := IsPasswordValid(passwd)
	if err != nil {
		response := FailedResponse{Error: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if isEmailValid == true && isPasswordValid == true {
		b := DoesUserExist(email)

		if b == true {
			response := FailedResponse{Error: "This user already exists"}

			json.NewEncoder(w).Encode(response)
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
	} else {
		fmt.Println(isEmailValid, isPasswordValid)
		response := FailedResponse{Error: "Email or password are not valid"}
		json.NewEncoder(w).Encode(response)
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("register.html")
	t.Execute(w, nil)
}
