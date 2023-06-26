package mirroredQuery

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type RespBody struct {
	text []byte
	err  error
}

func Run(servers [4]string) (text []byte, err error) {
	var wg sync.WaitGroup
	res := make(chan *RespBody)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case respb := <-res:
				if len(respb.text) > 0 {
					cancel()
					text = respb.text
					err = respb.err
					for range res {
					}
				}
			}
		}
	}()

	for i := range servers {
		srv := servers[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			httpDo(ctx, srv, res)
		}()
	}

	wg.Wait()
	log.Print("done")
	return
}

func httpDo(ctx context.Context, resource string, res chan<- *RespBody) {
	req, err := http.NewRequestWithContext(ctx, "GET", resource, nil)
	respb := new(RespBody)
	if err != nil {
		respb.err = fmt.Errorf("can't create request %s", err)
		res <- respb
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		respb.err = fmt.Errorf("can't perform request %s", err)
		res <- respb
	} else {
		text, err := io.ReadAll(resp.Body)
		if err != nil {
			respb.err = fmt.Errorf("can't read response %s", err)
			res <- respb
			return
		}
		respb.text = text
		res <- respb
	}
}
