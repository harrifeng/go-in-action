package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
	fmt.Println()
	os.Exit(0)
}
