package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := [...]int{1, 2, 2, 4, 5, 5, 6, 7, 8}
	s := []int{10, 11, 12, 13, 14, 15}
	ds := []string{"abc", "acb", "acb", "bca", "abc", "abc"}
	text := "I am\na\tdoctor"
	text = string(formatDoubleSpaces([]byte(text)))
	fmt.Printf("%s\n", text)
	fmt.Printf("%s\n", stringReverse([]byte(text)))
	ds = removeDouble(ds)
	fmt.Printf("%v\n", ds)
	reverse(a[:])
	reverse(s[2:])
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", s)
	reverse(s[:2])
	fmt.Printf("%d\n", s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func removeDouble(s []string) []string {
	cs := s[:0]
	for k, v := range s {
		if k < len(s)-1 && v != s[k+1] {
			cs = append(cs, v)
			fmt.Printf("double %d %s\n", k, v)
		}
		if k == len(s)-1 {
			cs = append(cs, v)
		}
	}
	return cs
}

func formatDoubleSpaces(s []byte) []byte {
	for k, v := range s {
		if unicode.IsSpace(rune(v)) {
			s[k] = ' '
		}
	}
	return s
}

func stringReverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
