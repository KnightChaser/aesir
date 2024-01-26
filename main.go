package main

import (
	"aesir/datastructure"
	"aesir/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/KnightChaser/sentinela"
	"github.com/tidwall/gjson"
)

type SysmonEvent struct {
	Event struct {
		EventData map[string]interface{} `json:"EventData"`
		System    struct {
			Channel       string            `json:"Channel"`
			Computer      string            `json:"Computer"`
			Correlation   map[string]string `json:"Correlation"`
			EventID       json.Number       `json:"EventID"`
			EventRecordID json.Number       `json:"EventRecordID"`
			Execution     struct {
				ProcessID json.Number `json:"ProcessID"`
				ThreadID  json.Number `json:"ThreadID"`
			} `json:"Execution"`
			Keywords string      `json:"Keywords"`
			Level    json.Number `json:"Level"`
			Opcode   json.Number `json:"Opcode"`
			Provider struct {
				Guid string `json:"Guid"`
				Name string `json:"Name"`
			} `json:"Provider"`
			Security struct {
				UserID string `json:"UserID"`
			} `json:"Security"`
			Task        json.Number `json:"Task"`
			TimeCreated struct {
				SystemTime time.Time `json:"SystemTime"`
			} `json:"TimeCreated"`
			Version json.Number `json:"Version"`
		} `json:"System"`
	} `json:"Event"`
}

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

	// Sysmon(System Monitor) log file in Windows
	stats, err := sentinela.ParseEVTX("D:/sampleEVTX.evtx")
	if err != nil {
		log.Fatal(err)
	}

	for _, stat := range stats.Event {

		// Jsonify Sysmon EVTX structure with proper data type
		// Because data.Event.EventData is different for every sysmon event ID,
		// Process EventData with the dedicated function.
		id := gjson.Get(stat, "Event.System.EventID").Int()
		data := SysmonEvent{}
		json.Unmarshal([]byte(stat), &data)
		data.Event.EventData = datastructure.EventDataStructureJsonify(id, stat)

		response, err := collection.InsertOne(context.TODO(), data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)

		fmt.Println("=========================================================================")
	}
}
