package main

import (
	"fmt"
	"os"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var s int
	i := 0
	values := make([]int, 0)
	for i <= 10 {
		if _, err := fmt.Scanln(&s); err != nil {
			fmt.Fprintf(os.Stderr, "Invalid char: %v", err)
		}
		values = append(values, s)
		i++
	}

	fmt.Printf("Sort: %d", sorting(values))
}

func sorting(val []int) []int {
	var root *tree
	for _, v := range val {
		root = add(root, v)
	}
	return appendValues(root, val[:0])
}

func appendValues(t *tree, values []int) []int {
	if t != nil {
		values = appendValues(t.left, values)
		values = append(values, t.value)
		values = appendValues(t.right, values)
	}
	return values
}

func add(t *tree, val int) *tree {
	if t == nil {
		t = new(tree)
		t.value = val
		return t
	}

	if val < t.value {
		t.left = add(t.left, val)
	} else {
		t.right = add(t.right, val)
	}
	return t
}
