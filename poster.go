package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const APIKEY = "7d625a7b"

type Poster struct {
	Title    string
	Year     string
	Writer   string
	Runtime  string `json:"runtime,omitempty"`
	Released string `json:"released,omitempty"`
	Rate     string `json:"imdbRating,omitempty"`
	Poster   string
}

func main() {
	title := strings.Join(os.Args[1:], " ")
	opt := url.Values{}
	opt.Add("t", title)
	opt.Add("apikey", APIKEY)

	p := send(opt.Encode())
	fmt.Printf("%s %s", p.Title, p.Poster)
}

func send(query string) *Poster {
	uri := "http://www.omdbapi.com?" + query
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	p := new(Poster)
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
