package webRequestHandler

import (
	"aesir/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Functions for API testing
func API(w http.ResponseWriter, r *http.Request) {
	// Create a map to represent the JSON data
	responseData := map[string]string{"data": "testing successful"}

	// Convert the map to JSON
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

// Get number of document in the given collection
func APISearchForDocumentCount(w http.ResponseWriter, r *http.Request) {
	// Receive API parameters
	apiParameters := mux.Vars(r)
	collectionInUse := apiParameters["collection"]

	// Set client options
	mongoDBURL := os.Getenv("DB_ACCESS_FULL_URL")
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)

	// Some additional code can be added here to perform operations with the MongoDB client
	// For example, you can use the 'client' variable to perform CRUD operations.
	dbname := os.Getenv("DB_NAME")
	db := client.Database(dbname)

	documentQty, err := db.Collection(collectionInUse).EstimatedDocumentCount(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	jsonResult := map[string]interface{}{
		"success":     true,
		"documentQty": documentQty,
	}

	responseJSON, _ := json.Marshal(jsonResult)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

// API for searching some objects with conditions in a particular MongoDB collection
// ex) Query db(aesir).C.find({"$and":[{"event.system.eventid": 3}, {"event.eventdata.DestinationIp": "8.8.8.8"}]})
//
//	=> localhost:8080/api/search/{mongodb_collection_name}/find/{mongodb_search_condition_in_JSON}
//	=> localhost:8080/api/search/C/find/%7B%22%24and%22%3A%5B%7B%22event.system.eventid%22%3A%203%7D%2C%20%7B%22event.eventdata.DestinationIp%22%3A%20%228.8.8.8%22%7D%5D%7D (URL encoded)
//
// Then it will return the JSONified result data after querying DB like an ordinary API
func APISearchWithCondition(w http.ResponseWriter, r *http.Request) {
	// Receive API parameters
	apiParameters := mux.Vars(r)
	collectionInUse := apiParameters["collection"]
	condition := apiParameters["condition"]
	request := apiParameters["request"]

	// Decode the condition string from the URL
	decodedCondition, err := url.QueryUnescape(condition)
	if err != nil {
		apiErrorJSONHandler(w, r, "Error decoding the search condition", http.StatusBadRequest)
		return
	}

	// Parse the decoded condition into a slice of BSON documents
	var searchCondition []bson.M
	err = bson.UnmarshalExtJSON([]byte(decodedCondition), true, &searchCondition)
	if err != nil {
		apiErrorJSONHandler(w, r, fmt.Sprintf("Error parsing the search condition: %v", err), http.StatusBadRequest)
		return
	}

	// Set client options
	mongoDBURL := os.Getenv("DB_ACCESS_FULL_URL")
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)
	collection := client.Database(os.Getenv("DB_NAME")).Collection(collectionInUse)

	// Perform the search using the specified condition
	// The condition will work the same with querying data to the MongoDB console directly; such like db.{collection}.find({"event.system.eventid": 3})
	var cursor *mongo.Cursor
	if request == "find" {
		cursor, err = collection.Find(context.Background(), searchCondition)
	} else if request == "aggregate" {
		cursor, err = collection.Aggregate(context.Background(), searchCondition)
	} else {
		apiErrorJSONHandler(w, r, fmt.Sprintf("Invalid request type, received %v", request), http.StatusBadRequest)
		return
	}
	defer cursor.Close(context.Background())

	// Iterate through the results and write them to the response
	var results []bson.M
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			apiErrorJSONHandler(w, r, fmt.Sprintf("Error decoding query results: %v", err), http.StatusInternalServerError)
			return
		}
		results = append(results, result)
	}

	// Convert results to JSON and write to the response
	jsonResults, err := json.Marshal(results)
	if err != nil {
		apiErrorJSONHandler(w, r, fmt.Sprintf("Error encoding query results to JSON: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResultsToReturn := map[string]interface{}{
		"success": true,
		"result":  string(jsonResults),
	}

	responseJSON, _ := json.Marshal(jsonResultsToReturn)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

// API for searching some objects with multiple conditions in a particular MongoDB collection
func APISearchWithMultipleCondition(w http.ResponseWriter, r *http.Request) {
	// Receive API parameters
	apiParameters := mux.Vars(r)
	collectionInUse := apiParameters["collection"]
	condition := apiParameters["condition"]

	// Decode the condition string from the URL
	decodedCondition, err := url.QueryUnescape(condition)
	if err != nil {
		apiErrorJSONHandler(w, r, "Error decoding the search condition", http.StatusBadRequest)
		return
	}

	var searchCondition map[string][]bson.M

	// Parse the decoded condition into a slice of BSON documents
	err = json.Unmarshal([]byte(decodedCondition), &searchCondition)

	if err != nil {
		apiErrorJSONHandler(w, r, fmt.Sprintf("Error parsing the search condition: %v", err), http.StatusBadRequest)
		return
	}

	// Set client options
	mongoDBURL := os.Getenv("DB_ACCESS_FULL_URL")
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)
	collection := client.Database(os.Getenv("DB_NAME")).Collection(collectionInUse)

	// Perform the search using the specified condition
	// The condition will work the same with querying data to the MongoDB console directly; such like db.{collection}.find({"event.system.eventid": 3})
	var cursor *mongo.Cursor
	searchOptions := options.Find().SetSort(bson.D{primitive.E{Key: "_id", Value: -1}})
	cursor, err = collection.Find(context.Background(), searchCondition, searchOptions)
	if err != nil {
		apiErrorJSONHandler(w, r, fmt.Sprintf("Error querying the database: %v", err), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	// Iterate through the results and write them to the response
	var results []bson.M
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			apiErrorJSONHandler(w, r, fmt.Sprintf("Error decoding query results: %v", err), http.StatusInternalServerError)
			return
		}
		results = append(results, result)
	}

	// Get the count of searched rows
	count, err := collection.CountDocuments(context.Background(), searchCondition)
	if err != nil {
		apiErrorJSONHandler(w, r, fmt.Sprintf("Error counting the number of documents: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert results to JSON and write to the response
	jsonResults, err := json.Marshal(results)
	if err != nil {
		apiErrorJSONHandler(w, r, fmt.Sprintf("Error encoding query results to JSON: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResultsToReturn := map[string]interface{}{
		"success": true,
		"result":  string(jsonResults),
		"count":   count,
	}

	responseJSON, _ := json.Marshal(jsonResultsToReturn)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
