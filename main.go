package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	store := InMemoryPlayerStore{}
	server := &PlayerServer{&store}
	log.Fatal(http.ListenAndServe(":5001", server))
}
