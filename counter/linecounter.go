package counter

import (
	"bufio"
	"bytes"
)

type LineCounter int

func (l *LineCounter) Write(p []byte) (n int, err error) {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	n = len(lines)
	*l += LineCounter(n)
	return
}
