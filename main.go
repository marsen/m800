package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	viper.SetDefault("mongo.url", "mongodb://localhost:27017")
	// Set client options
	clientOptions := options.Client().ApplyURI(viper.GetString("mongo.url"))

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

	// Get a handle for your collection
	collection := client.Database(viper.GetString("mongo.database")).Collection("homework")
	collection.Find(context.TODO(), bson.D{})
	fmt.Println("Connected homework")
	fmt.Println("Hello, World!")
}
