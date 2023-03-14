package main

import (
	"fmt"
	"math"
)

var z float64

func main() {
	for i := 0; i < 8; i++ {
		fmt.Printf("x = %d, e^x = %8.3f\n", i, math.Exp(float64(i)))
	}
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}
