package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var m = map[string]int{}

func main() {
	var s1, s2, s3 = []string{"Liza", "Kirill", "Olga"}, []string{"Kirill", "Liza", "Olga"}, []string{"Olga", "Kirill", "Liza"}
	add(s1)
	add(s2)
	add(s3)
	add(s2)
	for k, v := range m {
		fmt.Printf("%s\t%d\n", k, v)
	}
	ages := map[string]int{
		"Lisaveta": 16,
		"Olga":     49,
		"Vladimir": 52,
		"Kirill":   27,
	}
	ages2 := map[string]int{
		"Lisaveta": 16,
		"Olga":     49,
		"Vladimir": 52,
		"Kirill":   28,
	}
	dedup()
	fmt.Println(equal(ages2, ages))
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k := range x {
		if yv, ok := y[k]; !ok || x[k] != yv {
			return false
		}
	}
	return true
}

func dedup() {
	set := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if !set[text] {
			set[text] = true
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	for name := range set {
		fmt.Printf("*%s\n", name)
	}
}

func transform(s []string) string {
	return fmt.Sprintf("%q", s)
}

func add(s []string) {
	m[transform(s)]++
}

func all(s []string) int {
	return m[transform(s)]
}
