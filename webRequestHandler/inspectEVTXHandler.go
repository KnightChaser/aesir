package webRequestHandler

import (
	"aesir/db"
	"context"
	"fmt"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InspectEVTXHandler(w http.ResponseWriter, r *http.Request) {

	// Set client options
	mongoDBURL := "mongodb://localhost:27017"
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)

	// Some additional code can be added here to perform operations with the MongoDB client
	// For example, you can use the 'client' variable to perform CRUD operations.
	dbname := "aesir"
	db := client.Database(dbname)

	// Check if at least one EVTX collection exists
	collectionExists, err := checkEVTXCollectionExistence(db)
	if err != nil {
		fmt.Println("Error checking EVTX collection existence:", err)
		collectionExists = false
	}

	// If there is no uploaded EVTX collection on the database, immediately redirect to upload_sysmon_evtx.html.
	// If not (at least one EVTX collection exists), enlist collection and also provide a link to upload_sysmon_evtx.html
	if collectionExists {
		// Perform additional actions based on the existence of the collection, e.g., list collections

		// List all collections
		collections, err := listCollections(db)
		if err != nil {
			fmt.Println("Error listing collections:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Instruction2) Render welcome.html with the data containing every name of collection.
		// welcome.html is equivalent to the HTML file that you just created.
		data := struct {
			Collections []string
		}{
			Collections: collections,
		}

		// Render a template or provide additional information
		renderTemplate(w, "web/inspectEVTX/welcome.html", data)
	} else {
		// Redirect to upload_sysmon_evtx.html
		http.ServeFile(w, r, "web/inspectEVTX/upload_sysmon_evtx.html")
	}
}

func checkEVTXCollectionExistence(db *mongo.Database) (bool, error) {

	// Instruction1) If there is no collection in the database, it means there is no EVTX collection.
	// so, return true if one or more collections exist, false if not.

	collections, err := db.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return false, err
	}

	// Collection exists
	if len(collections) >= 1 {
		return true, nil
	}

	return false, nil
}

func listCollections(db *mongo.Database) ([]string, error) {
	// Retrieve a list of all collections in the database
	collections, err := db.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func renderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	// Load and render an HTML template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		// Handle the error, for simplicity, we'll log it
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		// Handle the error, for simplicity, we'll log it
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
