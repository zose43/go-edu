package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	words := make(map[string]int)
	f, e := os.Open(mustDoc())
	if e != nil {
		fmt.Fprintf(os.Stderr, "Can't open %s", e)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		words[s]++
	}
	fmt.Printf("Word\tCount\n")
	w := make([]string, 0)
	for word := range words {
		w = append(w, word)
	}
	sort.Strings(w)
	for _, n := range w {
		fmt.Printf("%s\t%d\n", n, words[n])
	}
}

func mustDoc() string {
	textdoc := flag.String(
		"f",
		"",
		"Textual document")
	flag.Parse()
	if *textdoc == "" {
		log.Fatal("No one doc given")
	}
	return *textdoc
}
