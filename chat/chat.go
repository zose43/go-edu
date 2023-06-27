package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func Run(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("start on %d", port)

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() { _ = conn.Close() }()
	ch := make(chan string)

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are: " + who
	entering <- ch
	messages <- "Connected " + who

	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		text := fmt.Sprintf("%s: %s", who, strings.TrimSpace(scan.Text()))
		messages <- text
	}

	leaving <- ch
	messages <- "Bye-bye " + who
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintf(conn, "%s\n", msg)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
