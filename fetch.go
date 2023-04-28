package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	fName := "files/urls"
	f, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	urls, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("%s %v\n", fName, err)
	}
	for _, u := range strings.Split(string(urls), "\n") {
		localF, size, err := fetching(u)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("url %s file %s size %d\n", u, localF, size)
		}
	}
}

func fetching(url string) (string, int64, error) {
	cl := http.DefaultClient
	cl.Timeout = time.Second * 5
	resp, err := cl.Get(url)
	if err != nil {
		return "---", 0, fmt.Errorf("%s", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("%s %s", url, err)
		}
	}()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create("files/" + local)
	if err != nil {
		return "---", 0, fmt.Errorf("%s not create path %s", url, local)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("%s %s", url, err)
		}
	}()

	s, err := io.Copy(f, resp.Body)
	if err != nil {
		return local, 0, fmt.Errorf("%s not copy content %s\t%v", url, local, err)
	}
	return local, s, nil
}
