package main

import (
	Tconv "./tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}
		f := Tconv.Fahrenheit(t)
		k := Tconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n", f, Tconv.FToC(f), k, Tconv.KToC(k))
	}
}
