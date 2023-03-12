package popCount

import "fmt"

// count bits where value not 0
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	r := byte(0)
	for i := 0; i <= 7; i++ {
		r += pc[byte(x>>(8*i))]
		fmt.Printf("index %d, value %d\n", i, r)
	}
	return int(r)
}
