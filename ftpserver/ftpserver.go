package ftpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func FtpServer(port int) {
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
		_, _ = conn.Write([]byte("Welcome to best FTP ever!\n"))
		handleCommands(conn)
	}
}

func handleCommands(conn net.Conn) {
	defer func() { _ = conn.Close() }()
	reader := bufio.NewReader(conn)
	for {
		cmdLine, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Can't read command %s", err)
		}

		cmd, arg, exist := strings.Cut(strings.TrimSpace(cmdLine), " ")
		if !exist && len(cmd) < 1 {
			log.Printf("Can't parse command %s", cmdLine)
			continue
		}

		switch cmd {
		case "ls":
			ls(arg)
		case "cd":
			cd(arg)
		case "get":
			get(arg)
		case "close":
			_, _ = conn.Write([]byte("Bye-bye"))
			_ = conn.Close()
		default:
			log.Printf("Unknown command %s", cmd)
		}
	}
}

func get(arg string) {

}

func cd(arg string) {

}

func ls(arg string) {

}
