package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.URL, r.Method, r.Proto)
	fmt.Fprintf(w, "Remote: %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Host: %q\n", r.Host)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]: %q\n", k, v)
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]: %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Requests count: %d\n", count)
	mu.Unlock()
}
