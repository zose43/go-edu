package http

import (
	"fmt"
	"net/http"
)

type dollars float32

type Database map[string]dollars

func (d Database) Show(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("i")
	price, ok := d[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Order: %q not found\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (d Database) List(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
