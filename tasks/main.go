package main

import (
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"
	log.Print("Running server on " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
