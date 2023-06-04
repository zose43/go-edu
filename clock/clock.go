package clock

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func Clock(port int, tFormat string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Start listen on port: %d\n", port)
	for {
		conn, err := listener.Accept()
		fmt.Printf("conn: %s\n", conn.RemoteAddr().String())
		if err != nil {
			log.Print(err)
			continue
		}
		go handle(conn, tFormat)
	}
}

func handle(conn net.Conn, tFormat string) {
	defer func() { _ = conn.Close() }()
	for {
		_, err := io.WriteString(conn, fmt.Sprintf("%s\n", time.Now().Format(tFormat)))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
