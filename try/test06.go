package main

import (
	"fmt"
	"net/http"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome2")
}

func main() {
	http.Handle("/", handler{})
	http.ListenAndServe(":8088", nil)
}
