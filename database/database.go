package database

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mu Mutex
var Mu sync.Mutex

// GetConnection returns mongodb client
func GetConnection() (*mongo.Client, error) {

	// Get the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Cancel the connection when we done
	defer cancel()

	// Get the client and handle error if any
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://zeref:12340987@api.38lef.mongodb.net/api?retryWrites=true&w=majority",
	))
	if err != nil {
		return nil, err
	}

	// Finally return client
	return client, nil
}

// GetCollection returns a collection from database
func GetCollection(DbName string, CollectionName string) (*mongo.Collection, error) {

	// Get the connection and handle error if any
	client, err := GetConnection()
	if err != nil {
		return nil, err
	}

	// Get the collection
	collection := client.Database(DbName).Collection(CollectionName)

	// Finally return the collection
	return collection, nil
}
