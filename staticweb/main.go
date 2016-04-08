package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development!")
}

func main() {

	http.HandleFunc("/welcome", messageHandler)

	log.Println("Listening...")

	http.ListenAndServe(":8080", nil)
	os.Exit(0)
}
