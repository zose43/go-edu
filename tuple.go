package main

import "fmt"

func main() {
	x, y := 156, 8
	fmt.Printf("Наибольший общий делитель: %d", gcd(x, y))
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
