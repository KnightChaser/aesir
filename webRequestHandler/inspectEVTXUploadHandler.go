package webRequestHandler

import (
	"aesir/datastructure"
	"aesir/db"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/KnightChaser/sentinela"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
)

func InspectEVTXUploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for the entire form
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, handler, err := r.FormFile("EVTXfile")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get the name from the form data
	dbCollectionName := r.FormValue("name")

	// Create a unique filename to avoid overwriting
	filename := filepath.Join("uploaded", handler.Filename)

	// Create the directory if it does not exist
	err = os.MkdirAll("uploaded", os.ModePerm)
	if err != nil {
		http.Error(w, "Unable to create directory", http.StatusInternalServerError)
		return
	}

	// Create a new file
	newFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the file content to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Unable to copy file content", http.StatusInternalServerError)
		return
	}

	// Set client options
	mongoDBURL := os.Getenv("DB_ACCESS_FULL_URL")
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)

	// Some additional code can be added here to perform operations with the MongoDB client
	// For example, you can use the 'client' variable to perform CRUD operations.
	dbname := os.Getenv("DB_NAME")
	db := client.Database(dbname)
	collection := db.Collection(dbCollectionName)

	// Sysmon(System Monitor) log file in Windows
	stats, err := sentinela.ParseEVTX(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, stat := range stats.Event {

		// Jsonify Sysmon EVTX structure with proper data type
		// Because data.Event.EventData is different for every sysmon event ID,
		// Process EventData with the dedicated function.
		id := gjson.Get(stat, "Event.System.EventID").Int()
		data := datastructure.SysmonEvent{}
		json.Unmarshal([]byte(stat), &data)
		data.Event.EventData = datastructure.EventDataStructureJsonify(id, stat)

		response, err := collection.InsertOne(context.TODO(), bson.M{"eventrecordid": data.Event.System.EventRecordID, "event": data.Event})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[+] EVTX event (%v) registered to %s/%s/%s\n", response, mongoDBURL, dbname, dbCollectionName)
	}

	// Redirect user to the page letting them know the file was uploaded successfully
	http.Redirect(w, r, "/web/entrance/upload_sysmon_evtx_successful.html", http.StatusSeeOther)

}
