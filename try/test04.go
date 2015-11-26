package main

import (
	"fmt"
	// "net/http"
)

type hfeng struct{}

func (h hfeng) Hello() {
	fmt.Println("Hello hfeng")
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Welcome")
// }

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	h := hfeng{}

	h.Hello()
}
