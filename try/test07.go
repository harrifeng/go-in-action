package main

import (
	"fmt"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v \n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are on about page")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome5")
}

func main() {
	commonHandlers := alice.New(loggingHandler)
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.ListenAndServe(":8080", nil)
}
