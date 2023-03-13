package main

import "fmt"

var x = 1<<1 | 1<<5
var y = 1<<1 | 1<<2

func main() {
	fmt.Printf("{1,5} %08b\n", x)
	fmt.Printf("{1,2} %08b\n", y)
	fmt.Printf("{1} %08b\n", x&y)     // intersection
	fmt.Printf("{1,2,5} %08b\n", x|y) // union, addiction if int = degree of 2
	fmt.Printf("{5,2} %08b\n", x^y)   // symmetric subtraction
	fmt.Printf("{5} %08b\n", x&^y)    // subtraction
	fmt.Printf("{5} %08b\n", ^x)      // inversion
	fmt.Printf("{3,7} %08b\n", x<<2)
	fmt.Printf("{0,4} %08b\n", x>>1)
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Printf("%d %#[1]o %#[1]X\n",i)
		}
	}
}
