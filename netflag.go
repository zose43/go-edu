package main

import (
	"fmt"
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMultiCast
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB // overhead int32 (1 << 32)
	PiB
	EiB
	ZiB // overhead in64 (1 << 64)
	Yib
)

const (
	KB = 1000
	MB = 1000 * 1000
	GB = 1000 * 1000 * 1000
	TB = 1000 * 1000 * 1000 * 1000
	PB = 1000 * 1000 * 1000 * 1000 * 1000
	EB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000
	ZB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000
	YB = 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000
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
