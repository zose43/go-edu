package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "аНаГрамма"
	sw := "ГраМма"
	fmt.Println(anagram(a, sw))
}

func anagram(a, sw string) bool {
	ang := make(map[rune]int8)
	a = strings.ToLower(a)
	sw = strings.ToLower(sw)
	for _, v := range a {
		ang[v]++
	}
	for _, v := range sw {
		if count, ok := ang[v]; ok && count > 0 {
			ang[v]--
		} else {
			return false
		}
	}
	return true
}
