package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	echoFromLoop()
}

func echoFromLoop() {
	for i, arg := range os.Args[1:] {
		fmt.Println(arg + "-" + strconv.Itoa(i))
	}
}

func echoAlternative() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
