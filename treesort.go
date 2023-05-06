package main

import (
	"bytes"
	"fmt"
	"os"
)

type Tree struct {
	value       int
	left, right *Tree
}

func (t *Tree) String() string {
	var res bytes.Buffer
	res.WriteByte('{')
	var preOrder func(node *Tree)
	preOrder = func(node *Tree) {
		if node == nil {
			return
		}
		if res.Len() > len("{") {
			res.WriteByte(' ')
		}
		res.WriteString(fmt.Sprintf("%d", node.value))
		preOrder(node.left)
		preOrder(node.right)
	}

	preOrder(t)
	res.WriteByte('}')
	return res.String()
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

	fmt.Printf("Sort: %d", Sorting(values))
}

func Sorting(val []int) []int {
	var root *Tree
	for _, v := range val {
		root = Add(root, v)
	}
	return AppendValues(root, val[:0])
}

func AppendValues(t *Tree, values []int) []int {
	if t != nil {
		values = AppendValues(t.left, values)
		values = append(values, t.value)
		values = AppendValues(t.right, values)
	}
	return values
}

func Add(t *Tree, val int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = val
		return t
	}

	if val < t.value {
		t.left = Add(t.left, val)
	} else {
		t.right = Add(t.right, val)
	}
	return t
}
