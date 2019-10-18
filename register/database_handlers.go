package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBUSER = os.Getenv("DBUSER")
	DBPASS = os.Getenv("DBPASS")
	DBSTR  = os.Getenv("DBSTR")
)

func InitConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://" + DBUSER + ":" + DBPASS + "@ds161710.mlab.com:61710/" + DBSTR + "?retryWrites=false")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func InsertUser(user User) {
	client := InitConnection()

	collection := client.Database(os.Getenv("DBSTR")).Collection("users")

	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a User: ", insertResult.InsertedID)
}

func DoesUserExist(email string) bool {
	filter := bson.D{{"email", email}}

	var result User

	client := InitConnection()

	collection := client.Database(os.Getenv("DBSTR")).Collection("users")

	collection.FindOne(context.TODO(), filter).Decode(&result)

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	if result.Email == email {
		return true
	}
	return false

}
