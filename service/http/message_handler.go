
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
    "net/http"
	"os"
)



func handleMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		// Handle the send message use case here
	case "GET":
		// get the unread messages here 

		// Get the whole chat history here basis on the request Path
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}