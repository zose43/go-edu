package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d", incr(&x))
}

func incr(p *int) int {
	*p++
	return *p
}
