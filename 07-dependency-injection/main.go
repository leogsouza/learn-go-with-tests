package main

import (
	"github.com/leogsouza/learn-go-with-tests/07-dependency-injection/greetings"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greetings.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(MyGreeterHandler))
}
