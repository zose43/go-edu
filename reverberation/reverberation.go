package reverberation

import (
	"bufio"
	"fmt"
	"log"
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
	signal := make(chan string)
	scanner := bufio.NewScanner(conn)

	fmt.Fprintf(conn, "Welcome, %s, to echo server\n", conn.RemoteAddr().String())

	go func() {
		defer func() {
			wg.Wait()
			_ = conn.Close()
		}()
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Fprintln(conn, "Timeout")
				return
			case phrase := <-signal:
				wg.Add(1)
				go func() {
					defer wg.Done()
					echo(conn, phrase, 1*time.Second)
				}()
			}
		}
	}()

	for scanner.Scan() {
		signal <- strings.TrimSpace(scanner.Text())
	}
}

func Start(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go Handle(conn)
	}
}
