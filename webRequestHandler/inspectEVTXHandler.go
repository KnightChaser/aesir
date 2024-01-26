package webRequestHandler

import "net/http"

func InspectEVTXHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the upload_sysmon_evtx.html file
	http.ServeFile(w, r, "web/inspectEVTX/upload_sysmon_evtx.html")
}
