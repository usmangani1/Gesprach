package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
    "net/http"
	"os"
)


func handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		// if path contains login handle the login feature

		// if the path contains logout handle the logout feature here.

		// Handle create user case here.
	case "GET":
		// get all the users here
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

