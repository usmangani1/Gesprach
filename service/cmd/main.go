package main

import (
	"log"
    "net/http"
)


func main() {
	http.HandleFunc("/order", handleOrders)
	log.Fatal(http.ListenAndServe(":8080", nil))
}