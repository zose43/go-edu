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
	defer close(done)

	c, err := sumFiles(done, root)
	if err := <-err; err != nil {
		return nil, err
	}

	for res := range c {
		if res.err != nil {
			log.Print(res.err)
			continue
		}
		m[res.path] = res.sum
	}

	return m, nil
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	ch := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			wg.Add(1)
			go func() {
				data, err := os.ReadFile(path)
				select {
				case ch <- result{path: path, sum: md5.Sum(data), err: err}:
				case <-done:
				}
				wg.Done()
			}()

			select {
			case <-done:
				return errors.New("cancelling")
			default:
				return nil
			}
		})

		go func() {
			wg.Wait()
			close(ch)
		}()
		errc <- err
	}()
	return ch, errc
}
