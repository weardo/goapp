package main

import (
	"log"
	"net/http"
)

func newInMemoryPlayserStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	server := &PlayerServer{newInMemoryPlayserStore()}
	log.Fatal(http.ListenAndServe(":5001", server))
}
