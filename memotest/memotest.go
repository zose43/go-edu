package memotest

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

var HTTPGetBody = httpGetBody

type M interface {
	Get(key string) (interface{}, error)
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()
	return io.ReadAll(resp.Body)
}

func incomingUrls() <-chan string {
	ch := make(chan string)
	go func() {
		for _, s := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io"} {
			ch <- s
		}
		close(ch)
	}()
	return ch
}

func Sequential(t *testing.T, m M) {
	for url := range incomingUrls() {
		start := time.Now()
		data, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s\t%s\t%.1fKB\n",
			url,
			time.Since(start),
			float64(len(data.([]byte))/1000))

	}
}

func Concurrent(t *testing.T, m M) {
	var wg sync.WaitGroup
	for url := range incomingUrls() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			data, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s\t%s\t%.1fKB\n",
				url,
				time.Since(start),
				float64(len(data.([]byte))/1000))
		}(url)
	}
	wg.Wait()
}
