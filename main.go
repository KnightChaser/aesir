package main

import (
	"aesir/datastructure"
	"aesir/db"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/KnightChaser/sentinela"
	"github.com/tidwall/gjson"
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
		data := datastructure.SysmonEvent{}
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
