package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupWithFiles(counts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dupWithFiles(counts map[string]int) {
	files := os.Args[1:]
	if len(files) > 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dumpWithFiles %v\n", f)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	} else {
		countLines(os.Stdin, counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
