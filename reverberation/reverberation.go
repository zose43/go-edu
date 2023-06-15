package reverberation

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
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
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		wg.Add(1)
		go func(phrase string) {
			defer wg.Done()
			echo(conn, phrase, 1*time.Second)
		}(scanner.Text())
	}
	defer func() {
		wg.Wait()
		_ = conn.Close()
	}()
}
