package main

import (
	"context"
	"fmt"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type PortfolioValue struct {
	Value float32
	Date  string
}

func DisconnectMongo(client *mongo.Client) {

	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func ConnectMongo() (client *mongo.Client) {
	// Set client options
	fmt.Println("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI("mongodb://172.22.128.1:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return
}

func InsertTransaction(client *mongo.Client, transaction DegiroTransaction) (id string) {
	collection := client.Database("portfolio-app").Collection("degiro-transactions")
	insertResult, err := collection.InsertOne(context.TODO(), transaction)
	if err != nil {
		log.Fatal(err)
	}

	id = fmt.Sprintf("%v", insertResult.InsertedID)
	fmt.Println("Inserted a single document: ", id)

	return
}

// func main() {

// 	client := connectMongo()
// 	mongoStuff(client)
// 	disconnectMongo(client)
// }
