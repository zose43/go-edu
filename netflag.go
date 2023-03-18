package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMultiCast
)

func main() {
	v := FlagMultiCast | FlagUp
	fmt.Printf("%b\t%t\n", v, isUp(v))
	turnDown(&v)
	fmt.Printf("%b\t%t\n", v, isUp(v))
	fmt.Printf("%b\t%t\n", v, isCast(v))
	setBroadcast(&v)
	fmt.Printf("%b\t%t\n", v, isUp(v))
	fmt.Printf("%b\t%t\n", v, isCast(v))
}

func isUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func turnDown(v *Flags) {
	*v &^= FlagUp
}

func setBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func isCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMultiCast) != 0
}
