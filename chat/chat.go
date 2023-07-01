package chat

import (
	"bufio"
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
	defer func() { _ = conn.Close() }()
	p := new(person)
	outer := make(chan string, 10)
	self := make(chan string)
	p.ch = outer

	go clientWriter(conn, outer)

	_, _ = fmt.Fprint(conn, "Write your name,pls\n")
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Print(err)
		p.name = conn.RemoteAddr().String()
	}
	p.name = strings.TrimSpace(name)

	outer <- "You are: " + p.name
	entering <- p
	messages <- "Connected " + p.name

	go func() {
		dur := 3 * time.Minute
		tick := time.NewTicker(dur)
		for {
			select {
			case <-tick.C:
				_, _ = fmt.Fprint(conn, "Connection close(timeout)")
				_ = conn.Close()
				return
			case <-self:
				tick.Reset(dur)
			}
		}
	}()

	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		text := fmt.Sprintf("%s: %s", p.name, strings.TrimSpace(scan.Text()))
		messages <- text
		self <- text
	}

	close(self)
	leaving <- p
	messages <- "Bye-bye " + p.name
}

func clientWriter(conn net.Conn, outer chan string) {
	for msg := range outer {
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
