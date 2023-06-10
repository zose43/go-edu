package netcat

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Netcat(port int, protocol string) {
	conn, err := net.Dial(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("Done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	_ = conn.(*net.TCPConn).CloseWrite()

	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
