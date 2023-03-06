package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	for _, url := range os.Args[1:] {
		match, _ := regexp.MatchString(`https?://`, url)
		if match != true {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s %v\n", url, err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "\nresponse-status: %s,%s\n", resp.Status, url)
	}
}
