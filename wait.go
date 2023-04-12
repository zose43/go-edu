package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		doc, err := waitForServer(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n%s", url, doc)
	}
}

func waitForServer(url string) ([]byte, error) {
	timeout := 1 * time.Minute
	deadline := time.Now().Add(timeout)
	cLog := log.New(os.Stdout, "wait: ", 0)
	for tries := 0; time.Now().Before(deadline); tries++ {
		resp, err := http.Get(url)
		if err == nil {
			t, _ := io.ReadAll(resp.Body)
			return t, err
		}
		cLog.Printf("Server %s no response, %d try, %v", url, tries, err)
		time.Sleep(time.Second << uint(tries))
	}
	return nil, fmt.Errorf("Server %s not avaialble, timeout %s\n", url, timeout)
}
