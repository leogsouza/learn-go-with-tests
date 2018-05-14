package main

import (
	"log"
	"net/http"

	"github.com/leogsouza/learn-go-with-tests/11-http-server/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)

	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
