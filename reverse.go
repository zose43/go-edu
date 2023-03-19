package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := []int{10, 11, 12, 13, 14, 15}
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
