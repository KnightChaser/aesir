package main

import (
	"aesir/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Setting up database
	mongoDBURL := "mongodb://localhost:27017"
	client := db.ConnectMongoDBSession(mongoDBURL)
	defer db.DisconnectMongoDBSession(client)

	// Setting up web server with gorilla/mux
	muxRouter := mux.NewRouter()

	muxRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	muxRouter.PathPrefix("/static/inspectEVTX").Handler(http.StripPrefix("/static/inspectEVTX", http.FileServer(http.Dir("./static/inspectEVTX"))))

	// Define a route for the root path "/"
	muxRouter.HandleFunc("/", homeHandler)
	muxRouter.HandleFunc("/inspectEVTX/", inspectEVTXHandler)

	// Start the server on port 8080
	listeningAddressPort := "0.0.0.0:8080"
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(listeningAddressPort, muxRouter)
}

// HomeHandler is the handler function for the home route "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the index.html file
	http.ServeFile(w, r, "static/index.html")
}

func inspectEVTXHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the upload_sysmon_evtx.html file
	http.ServeFile(w, r, "static/inspectEVTX/upload_sysmon_evtx.html")
}
