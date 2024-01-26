package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to the MongoDB
func ConnectMongoDBSession(accessURL string) *mongo.Client {

	// Grab access URL
	clientOptions := options.Client().ApplyURI(accessURL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Panic(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

	return client

}

// Disconnect the given mongoDB Session
func DisconnectMongoDBSession(mongoDBClientSession *mongo.Client) {
	if err := mongoDBClientSession.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}

}
