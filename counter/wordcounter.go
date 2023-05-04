package counter

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	var words []string
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	n = len(words)
	*w += WordCounter(n)
	return
}
