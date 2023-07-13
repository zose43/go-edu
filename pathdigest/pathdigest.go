package pathdigest

import (
	"crypto/md5"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	done := make(chan struct{})
	ch := make(chan result)
	defer close(done)

	path, err := walkDir(done, root)

	const workers = 20
	var wg = sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			digest(done, path, ch)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		if res.err != nil {
			log.Print(res.err)
			continue
		}
		m[res.path] = res.sum
	}

	if err := <-err; err != nil {
		return nil, err
	}
	return m, nil
}

func walkDir(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	ch := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(ch)
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			select {
			case ch <- path:
			case <-done:
				return errors.New("cancelling")
			}
			return nil
		})
		errc <- err
	}()
	return ch, errc
}

func digest(done <-chan struct{}, path <-chan string, ch chan<- result) {
	for p := range path {
		data, err := os.ReadFile(p)
		select {
		case <-done:
			return
		case ch <- result{sum: md5.Sum(data), path: p, err: err}:
		}
	}
}
