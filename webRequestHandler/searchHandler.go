package webRequestHandler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	apiParameters := mux.Vars(r)
	currentCollectionInUse := apiParameters["collection"]

	data := struct {
		Collection string
	}{
		Collection: currentCollectionInUse,
	}

	renderTemplate(w, "web/search/index.html", data)

}
