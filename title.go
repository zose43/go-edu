package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1:2][0]
	err := title(url)
	if err != nil {
		fmt.Println(err)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("no done %v", err)
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("Content-Type not a html/text")
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func visitNode(n *html.Node) {
	if n.Data == "title" && n.FirstChild != nil {
		fmt.Printf("Title: %s", n.FirstChild.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for el := n.FirstChild; el != nil; el = el.NextSibling {
		forEachNode(el, pre, post)
	}
	if post != nil {
		post(n)
	}
}
