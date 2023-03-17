package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	prefix := "sp-"
	suffix := ".jpeg"
	s := "sp-world"
	s1 := "world_and_peace.jpeg"
	fmt.Println(hasPrefix(s, prefix))
	fmt.Println(hasSuffix(s1, suffix))

	jps := "hello, ミスター"
	fmt.Printf("utf-8 %d, usually %d\n", utf8.RuneCountInString(jps), len(s))
	for i, r := range jps {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
