package main

import (
	"bytes"
)

func main() {
	s := "133555223"
	println(comma2(s))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	l := len(s)
	if l <= 3 {
		return s
	}
	for i := 0; i <= l-1; i++ {
		if (l-i)%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
