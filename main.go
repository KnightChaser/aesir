package main

import (
	"aesir/webRequestHandler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// ENV setting via godotenv pkg
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env file.")
	}

	// Setting up web server with gorilla/mux
	muxRouter := mux.NewRouter()

	// Serve static files
	staticSubrouter := muxRouter.PathPrefix("/web/").Subrouter()
	staticSubrouter.PathPrefix("/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))

	// Serve /web/entrance static files
	staticWebEntranceSubrouter := muxRouter.PathPrefix("/web/entrance/").Subrouter()
	staticWebEntranceSubrouter.PathPrefix("/").Handler(http.StripPrefix("/web/entrance/", http.FileServer(http.Dir("/web/entrance/"))))

	// Serve /web/inspect static files
	staticWebInspectSubrouter := muxRouter.PathPrefix("/web/inspect").Subrouter()
	staticWebInspectSubrouter.PathPrefix("/").Handler(http.StripPrefix("/web/inspect/", http.FileServer(http.Dir("/web/inspect/"))))

	// Serve font files
	assetSubrouter := muxRouter.PathPrefix("/web/asset/").Subrouter()
	assetSubrouter.PathPrefix("/font/").Handler(http.StripPrefix("/web/asset/font/", http.FileServer(http.Dir("/web/asset/font/"))))

	// Define main routes
	muxRouter.HandleFunc("/", webRequestHandler.HomeHandler)
	muxRouter.HandleFunc("/entrance", webRequestHandler.InspectEVTXHandler)
	muxRouter.HandleFunc("/entrance/upload", webRequestHandler.InspectEVTXUploadHandler)
	muxRouter.HandleFunc("/inspect", webRequestHandler.InspectEVTXAnalysisHandler)

	// Start the server on port 8080
	listeningAddressPort := "0.0.0.0:8080"
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(listeningAddressPort, muxRouter)
}
