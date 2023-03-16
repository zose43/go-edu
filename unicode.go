package main

import "fmt"

func main() {
	prefix := "sp-"
	suffix := ".jpeg"
	s := "sp-world"
	s1 := "world_and_peace.jpeg"
	fmt.Println(hasPrefix(s, prefix))
	fmt.Println(hasSuffix(s1, suffix))
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
