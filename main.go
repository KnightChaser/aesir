package main

import (
	"aesir/db"
	"fmt"
)

func main() {
	// Set client options
	mongoDBURL := "mongodb://localhost:27017"
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)

	// Some additional code can be added here to perform operations with the MongoDB client
	// For example, you can use the 'client' variable to perform CRUD operations.
	db := client.Database("aesir")
	collection := db.Collection("userB")
	fmt.Println(collection)
}
