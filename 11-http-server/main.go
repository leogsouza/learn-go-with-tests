package main

import (
	"log"
	"net/http"

	"github.com/leogsouza/learn-go-with-tests/11-http-server/server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {

	srv := &server.PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
