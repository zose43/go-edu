package main

import (
	"bytes"
	"fmt"
)

// alternative fmt.Sprint(values)
func main() {
	i := []int{21, 12, 1995}
	println(printInt(i))
}

func printInt(ints []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range ints {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, err := fmt.Fprintf(&buf, "%d", v)
		if err != nil {
			continue
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
