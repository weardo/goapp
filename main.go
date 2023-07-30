package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(newInMemoryPlayserStore())
	log.Fatal(http.ListenAndServe(":5001", server))
}
