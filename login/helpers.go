package main

import (
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ReturnError(w http.ResponseWriter, response FailedResponse) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, response)
}
