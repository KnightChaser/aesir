package main

import (
	"aesir/webRequestHandler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Setting up web server with gorilla/mux
	muxRouter := mux.NewRouter()

	// Serve static files
	staticSubrouter := muxRouter.PathPrefix("/web/").Subrouter()
	staticSubrouter.PathPrefix("/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))

	// Serve /web static files
	staticEntranceSubrouter := muxRouter.PathPrefix("/web/entrance/").Subrouter()
	staticEntranceSubrouter.PathPrefix("/").Handler(http.StripPrefix("/web/entrance/", http.FileServer(http.Dir("/web/entrance/"))))

	// Serve font files
	assetSubrouter := muxRouter.PathPrefix("/web/asset/").Subrouter()
	assetSubrouter.PathPrefix("/font/").Handler(http.StripPrefix("/web/asset/font/", http.FileServer(http.Dir("/web/asset/font/"))))

	// Define main routes
	muxRouter.HandleFunc("/", webRequestHandler.HomeHandler)
	muxRouter.HandleFunc("/entrance", webRequestHandler.InspectEVTXHandler)
	muxRouter.HandleFunc("/entrance/upload", webRequestHandler.InspectEVTXUploadHandler)

	// Start the server on port 8080
	listeningAddressPort := "0.0.0.0:8080"
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(listeningAddressPort, muxRouter)
}
