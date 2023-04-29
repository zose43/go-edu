package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	foo(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], true)
	os.Stdout.Write(buf[:n])
}

func foo(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	foo(x - 1)
}
