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

	muxRouter.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
	muxRouter.PathPrefix("/web/inspectEVTX/").Handler(http.StripPrefix("/web/inspectEVTX/", http.FileServer(http.Dir("/web/inspectEVTX/"))))
	muxRouter.PathPrefix("/web/asset/font/").Handler(http.StripPrefix("/web/asset/font/", http.FileServer(http.Dir("/web/asset/font/"))))

	// Define a route for the root path "/"
	muxRouter.HandleFunc("/", webRequestHandler.HomeHandler)
	muxRouter.HandleFunc("/inspectEVTX/", webRequestHandler.InspectEVTXHandler)
	muxRouter.HandleFunc("/inspectEVTX/upload", webRequestHandler.InspectEVTXUploadHandler)

	// Start the server on port 8080
	listeningAddressPort := "0.0.0.0:8080"
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(listeningAddressPort, muxRouter)
}
