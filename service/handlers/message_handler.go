package handlers

import (
	"net/http"
	"strings"
)

func handleMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		// Handle the send message use case here
		if strings.Contains(r.URL.Path, "send/text") {

		}

	case "GET":
		// get the unread messages here
		if strings.Contains(r.URL.Path, "unread") {

		} else if strings.Contains(r.URL.Path, "history") {
			// handle the history here.
		}

		// Get the whole chat history here basis on the request Path
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
