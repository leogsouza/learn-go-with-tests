package main

import (
	"log"
	"net/http"

	"github.com/leogsouza/learn-go-with-tests/11-http-server/server"
)

func main() {

	srv := &server.PlayerServer{server.NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
