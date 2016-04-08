package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go web Development")
}

func main() {
	mux := http.NewServeMux()

	mh := http.HandlerFunc(messageHandler)

	mux.Handle("/welcome", mh)

	log.Println("Listening...")

	http.ListenAndServe(":8080", mux)
	os.Exit(0)
}
