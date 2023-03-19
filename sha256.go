package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

var bc [256]byte

func init() {
	for i := range bc {
		bc[i] = bc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	c1bits := countBits(c1)
	c2bits := countBits(c2)
	fmt.Printf("Bits diff: %d\n", int(math.Abs(float64(c1bits))-float64(c2bits)))
}

func countBits(sha [sha256.Size]byte) uint {
	r := byte(0)
	for _, v := range sha {
		r += bc[v]
	}
	return uint(r)
}
