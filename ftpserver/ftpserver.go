package ftpserver

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
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
		fmt.Fprintln(conn, "Welcome to best FTP ever!")
		go handleCommands(conn)
	}
}

func handleCommands(conn net.Conn) {
	defer func() { _ = conn.Close() }()
	reader := bufio.NewReader(conn)

	for {
		cmdLine, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Can't read command %s", err)
			return
		}

		cmd, arg, exist := strings.Cut(strings.TrimSpace(cmdLine), " ")
		if !exist && len(cmd) < 1 {
			log.Printf("Can't parse command %s", cmdLine)
			continue
		}

		switch cmd {
		case "ls":
			if len(arg) == 0 {
				arg = "."
			}
			execute(cmd, conn, arg)
		case "cd":
			// todo realize
			//execute("bash", conn, "-c", cmdLine)
		case "get":
			execute("cat", conn, arg)
		case "close":
			fmt.Fprintln(conn, "Bye-bye")
			_ = conn.Close()
		default:
			fmt.Fprintf(conn, "unknown command %s\n", cmd)
		}
	}
}

func execute(cmdLine string, conn io.Writer, arg ...string) {
	cmd := exec.Command(cmdLine, arg...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Print(err)
		fmt.Fprintf(conn, "failure execute %s, please try later\n", cmd)
	}
	_, _ = conn.Write(res)
}
