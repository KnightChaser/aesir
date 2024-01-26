package datastructure

import (
	"encoding/json"
	"time"
)

type SysmonEvent struct {
	Event struct {

		// EventData field varies according to Sysmon Event ID.
		// Refer to the datastructure/eventDataStructureJsonify.go for detailed specification
		EventData map[string]interface{} `json:"EventData"`

		System struct {
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
