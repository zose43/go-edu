package crawl

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

var tokens = make(chan struct{}, 20)

func Crawl(url, mainUrl string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := extract(url, mainUrl)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func extract(source, mainUrl string) ([]string, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Close }()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("can't get %s, response status %s", source, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var urls []string
	var findUrls func(node *html.Node)
	findUrls = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					u, err := resp.Request.URL.Parse(attr.Val)
					if err != nil {
						log.Print(err)
						break
					}
					if strings.Contains(u.String(), mainUrl) {
						urls = append(urls, u.String())
					}
				}
			}
		}
	}
	findElements(doc, findUrls)
	return urls, nil
}

func findElements(n *html.Node, f func(node *html.Node)) {
	if f != nil {
		f(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findElements(c, f)
	}
}

func Run(depth int) {
	worklist := make(chan []string)
	seen := make(map[string]bool)
	main := os.Args[1]
	var num int
	nextLevel := -1
	currentLevel := 1

	num++
	go func() { worklist <- os.Args[1:2] }()

	for ; num > 0; num-- {
		list := <-worklist
		if depth == 0 {
			break
		}
		nextLevel += len(list)
		if currentLevel == 0 {
			depth--
			currentLevel = nextLevel
		}
		currentLevel--
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				num++
				go func(u string) {
					worklist <- Crawl(u, main)
				}(link)
			}
		}
	}
	fmt.Printf("Done %d links", len(seen))
}
