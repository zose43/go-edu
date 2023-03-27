package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"realised"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var titles []struct{ Title string }
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newan"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("Invalid marshaling %s", err)
	}
	fmt.Printf("%s\n", data)

	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Invalid unmarshaling %s", err)
	}
	fmt.Printf("%s\n", titles)
}
