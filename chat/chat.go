package chat

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client chan<- string

type person struct {
	name string
	ch   client
}

var (
	entering = make(chan *person)
	leaving  = make(chan *person)
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		_ = conn.Close()
		cancel()
	}()
	p := new(person)
	ch := make(chan string)
	p.ch = ch

	go clientWriter(conn, ch)

	_, _ = fmt.Fprint(conn, "Write your name,pls\n")
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Print(err)
		p.name = conn.RemoteAddr().String()
	}
	p.name = strings.TrimSpace(name)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			ch <- "You are: " + p.name
			entering <- p
			messages <- "Connected " + p.name
			scan := bufio.NewScanner(conn)
			for scan.Scan() {
				text := fmt.Sprintf("%s: %s", p.name, strings.TrimSpace(scan.Text()))
				messages <- text
			}
		}
	}
	leaving <- p
	messages <- "Bye-bye " + p.name
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintf(conn, "%s\n", msg)
	}
}

func broadcaster() {
	clients := make(map[*person]bool)
	for {
		select {
		case msg := <-messages:
			for prs := range clients {
				prs.ch <- msg
			}
		case prs := <-entering:
			clients[prs] = true
			prs.ch <- fmt.Sprintf("Online(%d):", len(clients))
			for cli := range clients {
				prs.ch <- cli.name
			}
		case prs := <-leaving:
			delete(clients, prs)
			close(prs.ch)
		}
	}
}
