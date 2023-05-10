package http

import (
	"fmt"
	"net/http"
)

type dollars float32

type Database map[string]dollars

func (d Database) show(w http.ResponseWriter, item string) {
	if price, ok := d[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Order: %q not fount\n", item)
	}
}

func (d Database) list(w http.ResponseWriter) {
	for item, price := range d {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (d Database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		d.list(w)
	case "/price":
		qv := r.URL.Query().Get("i")
		d.show(w, qv)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page %s not found\n", r.URL)
	}
}
