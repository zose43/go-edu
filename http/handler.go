package http

import (
	"fmt"
	"net/http"
	"strconv"
)

type dollars float32

type Database map[string]dollars

func (d Database) Update(w http.ResponseWriter, r *http.Request) {
	qvalues := r.URL.Query()
	if !qvalues.Has("i") || !qvalues.Has("p") {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can't create item, empty params\n")
		return
	}

	item := qvalues.Get("i")
	price, err := strconv.ParseFloat(qvalues.Get("p"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can't create item, invalid price %q\n", qvalues.Get("p"))
		return
	}
	if _, ok := d[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Can't create item,%q not exist\n", item)
		return
	}

	d[item] = dollars(price)
	http.Redirect(w, r, fmt.Sprintf("/show?i=%s", item), 302)
}

func (d Database) Delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("i")
	delete(d, item)
	http.Redirect(w, r, "/list", 302)
}

func (d Database) Create(w http.ResponseWriter, r *http.Request) {
	qvalues := r.URL.Query()
	if !qvalues.Has("i") || !qvalues.Has("p") {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can't create item, empty params\n")
		return
	}

	item := qvalues.Get("i")
	price, err := strconv.ParseFloat(qvalues.Get("p"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can't create item, invalid price %q\n", qvalues.Get("p"))
		return
	}
	if _, ok := d[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can't create item, double %q\n", item)
		return
	}

	d[item] = dollars(price)
	http.Redirect(w, r, fmt.Sprintf("/show?i=%s", item), 302)
}

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
