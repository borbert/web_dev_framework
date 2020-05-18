package database

import (
"context"
"fmt"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"log"
"time"
)

var DB *mongo.Database

func InitDB(){
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://appDev:appPass@testapp-5qjs9.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to MongoDB!")

	DB = client.Database("go-web-dev-db")

}
