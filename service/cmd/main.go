package main

import (
	"log"
	"net/http"
)

func main() {
	handlers "github.com/usmangani1/gesprach/service/handlers"
	http.HandleFunc("/user", handlers.handleUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
