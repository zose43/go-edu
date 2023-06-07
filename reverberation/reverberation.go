package reverberation

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, phrase string, dur time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(phrase))
	time.Sleep(dur)
	fmt.Fprintln(c, "\t", phrase)
	time.Sleep(dur)
	fmt.Fprintln(c, "\t", strings.ToLower(phrase))
}

func Handle(conn net.Conn) {
	defer func() { _ = conn.Close() }()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		go echo(conn, scanner.Text(), 1*time.Second)
	}
}
