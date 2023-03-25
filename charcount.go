package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var utype = make(map[string]int)

func main() {
	count := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, s, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount %v\n", err)
			os.Exit(1)
		}
		if s == 1 && r == unicode.ReplacementChar {
			invalid++
			continue
		}
		defineType(r)
		count[r]++
		utflen[s]++
	}
	fmt.Printf("rune\tcount\n")
	for r, count := range count {
		fmt.Printf("%q\t%d\n", r, count)
	}
	for n, count := range utflen {
		if n > 0 {
			fmt.Printf("len=%d\tcount=%d\n", n, count)
		}
	}
	for t, c := range utype {
		fmt.Printf("%s\t%d\n", t, c)
	}
	if invalid > 0 {
		fmt.Printf("Invalid utf-8 symbols: %d\n", invalid)
	}
}

func defineType(r rune) {
	if unicode.IsLetter(r) {
		utype["Letter"]++
	}
	if unicode.IsNumber(r) {
		utype["Number"]++
	}
	if unicode.IsSpace(r) {
		utype["Space"]++
	}
	if unicode.IsPunct(r) {
		utype["Punctuational"]++
	}
}
