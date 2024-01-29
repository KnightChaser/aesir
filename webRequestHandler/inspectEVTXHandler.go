package webRequestHandler

import (
	"aesir/db"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InspectEVTXHandler(w http.ResponseWriter, r *http.Request) {

	// Set client options
	mongoDBURL := os.Getenv("DB_ACCESS_FULL_URL")
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)

	// Some additional code can be added here to perform operations with the MongoDB client
	// For example, you can use the 'client' variable to perform CRUD operations.
	dbname := os.Getenv("DB_NAME")
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

		// List all collections with additional information
		collectionsInfo, err := listCollectionsInfo(db)
		if err != nil {
			fmt.Println("Error listing collections:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Instruction2) Render welcome.html with the data containing every name of collection.
		// welcome.html is equivalent to the HTML file that you just created.
		data := struct {
			CollectionsInfo []CollectionInfo
		}{
			CollectionsInfo: collectionsInfo,
		}

		// Render a template or provide additional information
		renderTemplate(w, "web/entrance/welcome.html", data)
	} else {
		// Redirect to upload_sysmon_evtx.html
		http.ServeFile(w, r, "web/entrance/upload_sysmon_evtx.html")
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
	return len(collections) >= 1, nil
}

// CollectionInfo represents information about a MongoDB collection
type CollectionInfo struct {
	Name        string
	DocumentQty int64
	CreatedTime time.Time
	URL         string
}

func listCollectionsInfo(db *mongo.Database) ([]CollectionInfo, error) {
	// Retrieve a list of all collections in the database with additional information
	collections, err := db.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var collectionsInfo []CollectionInfo

	// Prepare every collection's name, # of documents, and creation time.
	// Creation time can be estimated by the first document in the collection's objectId which has a timestamp information.
	// MongoDB provides ObjectId.getTimestamp()
	for _, collectionName := range collections {
		// Count documents in the collection
		documentQty, err := db.Collection(collectionName).EstimatedDocumentCount(context.Background())
		if err != nil {
			return nil, err
		}

		// Get the first document's ObjectId to calculate the creation time
		var firstDoc map[string]interface{}
		err = db.Collection(collectionName).FindOne(context.Background(), bson.M{}).Decode(&firstDoc)
		if err != nil {
			return nil, err
		}

		// Extract and convert the ObjectId to a timestamp
		objectID, ok := firstDoc["_id"].(primitive.ObjectID)
		if !ok {
			return nil, fmt.Errorf("unable to extract ObjectID from the document")
		}
		createdTime := objectID.Timestamp()

		interactivePageAccessURL := fmt.Sprintf("/inspect/%s", collectionName)
		// Append collection info to the list
		info := CollectionInfo{
			Name:        collectionName,
			DocumentQty: documentQty,
			CreatedTime: createdTime,
			URL:         interactivePageAccessURL,
		}
		collectionsInfo = append(collectionsInfo, info)
	}

	return collectionsInfo, nil
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
