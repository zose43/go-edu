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
	defer func() {
		_ = conn.Close()
	}()

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}