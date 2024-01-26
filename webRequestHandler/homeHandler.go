package webRequestHandler

import "net/http"

// HomeHandler is the handler function for the home route "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the index.html file
	http.ServeFile(w, r, "web/index.html")
}
