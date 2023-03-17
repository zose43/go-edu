package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println(basename2(mustPath()))
}

func mustPath() string {
	p := flag.String("p", "", "path")
	flag.Parse()
	if *p == "" {
		log.Fatalf("No one arg %v", *p)
	}
	return *p
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // if no find, return -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return s[:dot]
	}
	return s
}
