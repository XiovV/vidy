package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	passwd := r.FormValue("password")

	client := InitConnection()

	collection := client.Database("mernstackcourse").Collection("users")

	newUser := User{email, passwd}

	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

	fmt.Printf("%s %s", email, passwd)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("register.html")
	t.Execute(w, nil)
}
