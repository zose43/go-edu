package ftpserver

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const FtpAbsolutePath = "/home/zose/go_projects/go-edu/files/ftp"

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
		printToClient("Welcome to best FTP ever!", conn)
		go handleCommands(conn)
	}
}

func handleCommands(conn net.Conn) {
	defer func() { _ = conn.(*net.TCPConn).CloseWrite() }()
	reader := bufio.NewReader(conn)

	for {
		cmdLine, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Can't read command %s", err)
			printToClient("Can't read command, please try later", conn)
			return
		}

		cmdArgs := strings.Fields(strings.TrimSpace(cmdLine))
		if len(cmdArgs) == 0 {
			printToClient("Empty request", conn)
			return
		}

		arg := ""
		if len(cmdArgs) > 1 {
			arg = cmdArgs[1]
		}

		switch cmdArgs[0] {
		case "list":
			list(arg, conn)
		case "cd":
			changeDir(arg, conn)
		case "get":
			get(arg, conn)
		case "close":
			printToClient("Bye-bye", conn)
			_ = conn.Close()
		default:
			printToClient(fmt.Sprintf("Unknown command %q", cmdArgs[0]), conn)
		}
	}
}

func path(relative string) (*os.File, error) {
	dir, _ := os.Getwd()
	openDir := fmt.Sprintf("files/ftp/%s", relative)
	if strings.Contains(dir, "/ftp") {
		openDir = fmt.Sprintf("%s/%s", dir, relative)
	}
	if relative == FtpAbsolutePath {
		openDir = relative
	}

	openDir = strings.TrimRight(openDir, "/")
	f, err := os.Open(openDir)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func list(arg string, conn net.Conn) {
	msg := ""
	f, err := path(arg)
	if err != nil {
		msg = fmt.Sprintf("Can't open path %q", arg)
		printToClient(msg, conn)
		return
	}

	files, err := f.Readdir(0)
	if err != nil {
		msg = fmt.Sprintf("Can't read directory %q", arg)
		printToClient(msg, conn)
		return
	}

	for _, f := range files {
		fType := "file"
		if f.IsDir() {
			fType = "dir"
		}
		_, _ = fmt.Fprintf(conn, "%s %s %d %s %s\n",
			f.Mode().Perm(),
			fType,
			f.Size(),
			f.ModTime().Format(time.DateTime),
			f.Name())
	}
}

func changeDir(arg string, conn net.Conn) {
	msg := ""
	if len(arg) == 0 {
		arg = FtpAbsolutePath
	}
	f, err := path(arg)
	if err != nil {
		msg = fmt.Sprintf("Can't open dir %q", arg)
		printToClient(msg, conn)
		return
	}
	if err := f.Chdir(); err != nil {
		msg = fmt.Sprintf("Can't change dir %q", arg)
		printToClient(msg, conn)
	}
}

func get(arg string, conn net.Conn) {
	msg := ""
	if len(arg) == 0 {
		printToClient("Filename is empty", conn)
		return
	}

	f, err := path(arg)
	if err != nil {
		msg = fmt.Sprintf("Can't open file %q", arg)
		printToClient(msg, conn)
		return
	}

	info, err := f.Stat()
	if err != nil {
		msg = fmt.Sprintf("Can't read file %q", arg)
		printToClient(msg, conn)
		return
	}
	if info.IsDir() {
		printToClient("Read only from dir", conn)
		return
	}

	text, err := io.ReadAll(f)
	if err != nil {
		msg = fmt.Sprintf("Can't read file %q", arg)
		printToClient(msg, conn)
		return
	}
	printToClient(string(text), conn)
}

func printToClient(msg string, conn net.Conn) {
	_, _ = fmt.Fprintln(conn, msg)
}
