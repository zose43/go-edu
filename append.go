package main

import (
	"fmt"
)

func main() {
	var y, x []int
	for i := 0; i <= 10; i++ {
		if i&1 != 0 {
			y = appendInt(y, i, i*i)
		} else {
			y = appendInt(y, i)
		}
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	}
	x = append(x, 2007, 2006)
	fmt.Printf("len=%d cap=%d\t%v\n", len(x), cap(x), x)
	bd := []int{21, 12, 1995}
	x = append(x, bd...)
	fmt.Printf("len=%d cap=%d\t%v\n", len(x), cap(x), x)
}

func appendInt(s []int, x ...int) []int {
	var z []int
	zlen := len(s) + len(x)
	if zlen <= cap(s) {
		z = s[:zlen]
	} else {
		zcap := zlen
		if zcap < len(s)*2 {
			zcap = len(s) * 2
		}
		z = make([]int, zlen, zcap)
		copy(z, s)
	}
	copy(z[len(s):], x)
	return z
}
