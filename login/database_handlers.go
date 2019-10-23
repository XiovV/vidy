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
	DEV    = os.Getenv("DEV")
)

func InitConnection() *mongo.Client {
	var clientOptions *options.ClientOptions

	fmt.Println(DBUSER, DBPASS, DBSTR)
	if DEV == "true" {
		clientOptions = options.Client().ApplyURI("mongodb://" + DBUSER + ":" + DBPASS + "@ds251894.mlab.com:51894/" + DBSTR + "?retryWrites=false")
	} else {
		clientOptions = options.Client().ApplyURI("mongodb://" + DBUSER + ":" + DBPASS + "@ds161710.mlab.com:61710/" + DBSTR + "?retryWrites=false")
	}
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

func CheckCredentials(email, password string) (bool, error) {
	filter := bson.D{{"email", email}}

	var result User

	client := InitConnection()

	collection := client.Database(DBSTR).Collection("users")

	collection.FindOne(context.TODO(), filter).Decode(&result)

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	isHashCorrect := CheckPasswordHash(password, result.Password)

	if result.Email == email && isHashCorrect == true {
		return true, nil
	}
	return false, fmt.Errorf("Email or password is incorrect")
}

func DoesUserExist(email string) bool {
	filter := bson.D{{"email", email}}

	var result User

	client := InitConnection()

	collection := client.Database(DBSTR).Collection("users")

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
