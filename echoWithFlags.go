package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Пропуск символа новой строки")
var s = flag.String("s", " ", "Разделитель")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}
