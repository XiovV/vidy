package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

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

	fmt.Println(reflect.TypeOf(client))

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
