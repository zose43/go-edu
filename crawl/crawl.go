package crawl

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var tokens = make(chan struct{}, 20)

func Crawl(ctx context.Context, url, mainUrl string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := extract(ctx, url, mainUrl)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func extract(ctx context.Context, source, mainUrl string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", source, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
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

func Run(depth int, main string) {
	var (
		worklist       = make(chan []string)
		seen           = make(map[string]bool)
		num, interrupt int
		nextLevel      = -1
		currentLevel   = 1
		wg             = sync.WaitGroup{}
	)

	num++
	go func() { worklist <- os.Args[1:2] }()

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	go func() {
		_, _ = os.Stdout.Read(make([]byte, 1))
		cancelFunc()
		interrupt = 1
	}()

	go func() {
		select {
		case <-ctx.Done():
			log.Printf("stop crawler: %v", ctx.Err())
			for range worklist {
			}
		}
	}()

	for ; num > 0; num-- {
		list := <-worklist
		if depth == 0 || interrupt == 1 {
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
				wg.Add(1)
				go func(u string) {
					defer wg.Done()
					worklist <- Crawl(ctx, u, main)
				}(link)
			}
		}
	}

	wg.Wait()
	fmt.Printf("Done %d links", len(seen))
}
