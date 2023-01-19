package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"m800/interal/dto"
)

var client *mongo.Client

type MongoImpl struct {
	c mongo.Client
}

func NewMongoImpl() MongoImpl {

	mongoClient := getClient(viper.GetString("mongo.url"))
	return MongoImpl{
		c: *mongoClient,
	}
}

func getClient(uri string) *mongo.Client {
	if client == nil {
		// Set client options
		clientOptions := options.Client().ApplyURI(uri)

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
		return client
	}
	return client

}

func (m *MongoImpl) Save(msg dto.Message) {

	// Get a handle for your collection
	collection := m.c.Database("test").Collection("messages")
	// Insert message into MongoDB
	insertResult, err := collection.InsertOne(context.TODO(), msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inserted a message: ", insertResult.InsertedID)
}
