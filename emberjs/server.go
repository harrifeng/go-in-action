package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Start Server")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
