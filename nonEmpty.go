package main

import "fmt"

func main() {
	months := []string{"January", "March", "April", "May", "", "July", "", "", "October", "", "December"}
	nonEmpty := nonEmpty2(months)
	fmt.Printf("%v\n", nonEmpty)
	fmt.Printf("%v\n", months)
	fmt.Printf("%v\n", remove(nonEmpty, 4))
}

func nonEmpty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

func nonEmpty2(s []string) []string {
	out := s[:0]
	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func remove(s []string, x int) []string {
	copy(s[x:], s[x+1:])
	return s[:len(s)-1]
}

func remove2(slice []string, x int) []string {
	slice[x] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
