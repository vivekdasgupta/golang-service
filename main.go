package main

import (
	"log"
	"net/http"
)

const PROG = "drmshows"
const PORT = ":5000"

func main() {

	router := NewRouter()

	log.Printf("%s\t starting on port %s\n", PROG, PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
