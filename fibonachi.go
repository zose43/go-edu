package main

import "fmt"

func main() {
	fmt.Printf("Fibonachi(n-1): %d", fib(12))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, y+x
	}
	return x
}
