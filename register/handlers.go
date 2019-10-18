package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	passwd := r.FormValue("password")

	hashedPasswd, err := HashPassword(passwd)

	client := InitConnection()

	collection := client.Database(os.Getenv("DBSTR")).Collection("users")

	newUser := User{email, hashedPasswd}

	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a User: ", insertResult.InsertedID)

	fmt.Println("tokenString", GenerateToken(newUser.Email))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("register.html")
	t.Execute(w, nil)
}
