package main

import (
	"aesir/webRequestHandler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func isDockerEnvironment() bool {
	// Check if the DOCKER_RUN environment variable is set
	_, isDocker := os.LookupEnv("DOCKER_RUN")
	return isDocker
}

func main() {

	// ENV setting via godotenv pkg
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env file.")
	}

	// Set the environment variable for the database URL if the application is not running in a Docker container
	// For local development, the environment variable is set in the .env file
	if !isDockerEnvironment() {
		os.Setenv("DB_ACCESS_FULL_URL", "mongodb://localhost:27017")
	}

	fmt.Printf("Database will be accessible to %s\n", os.Getenv("DB_ACCESS_FULL_URL"))

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

	// Serve /web/search sftatic files
	staticWebSearchSubrouter := muxRouter.PathPrefix("/web/search").Subrouter()
	staticWebSearchSubrouter.PathPrefix("/").Handler(http.StripPrefix("/web/search/", http.FileServer(http.Dir("/web/search/"))))

	// Serve font files
	assetSubrouter := muxRouter.PathPrefix("/web/asset/").Subrouter()
	assetSubrouter.PathPrefix("/font/").Handler(http.StripPrefix("/web/asset/font/", http.FileServer(http.Dir("/web/asset/font/"))))

	// Define main routes
	muxRouter.HandleFunc("/", webRequestHandler.HomeHandler)
	muxRouter.HandleFunc("/entrance", webRequestHandler.InspectEVTXHandler)
	muxRouter.HandleFunc("/entrance/upload", webRequestHandler.InspectEVTXUploadHandler)
	muxRouter.HandleFunc("/inspect/{collection}", webRequestHandler.InspectEVTXAnalysisHandler)
	muxRouter.HandleFunc("/search/{collection}", webRequestHandler.SearchHandler)
	muxRouter.HandleFunc("/api", webRequestHandler.API) // default
	muxRouter.HandleFunc("/api/search/{collection}/documentCount", webRequestHandler.APISearchForDocumentCount)
	muxRouter.HandleFunc("/api/search/{collection}/{request}/{condition}", webRequestHandler.APISearchWithCondition)
	muxRouter.HandleFunc("/api/searchMultipleCondition/{collection}/{condition}", webRequestHandler.APISearchWithMultipleCondition)

	// Start the server on port 8080
	listeningAddressPort := "0.0.0.0:8080"
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(listeningAddressPort, muxRouter)
}
