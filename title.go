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

	if err = soleTitle(doc); err != nil {
		return fmt.Errorf("url:%s %s", url, err)
	}
	return nil
}

func soleTitle(n *html.Node) (err error) {
	type bailout struct{}
	var title string
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("more than 1 title on page")
		default:
			panic(p)
		}
	}()

	forEachNode(n, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
			fmt.Printf("Title: %s", title)
		}
	}, nil)
	return nil
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
