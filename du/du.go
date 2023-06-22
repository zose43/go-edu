package du

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 10)

type PathInfo struct {
	Size      int64
	Name      string
	FileCount int64
}

func (pi *PathInfo) walkDir(dir string, filesize chan<- *PathInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go pi.walkDir(subdir, filesize, wg)
		} else {
			stat, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "can't get info about %s %v\n", entry.Name(), err)
			}
			pi.Size += stat.Size()
			pi.FileCount++
			filesize <- pi
		}
	}
}

func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := os.ReadDir(dir)

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read %s %v\n", dir, err)
		return nil
	}
	return entries
}

func ScanDirs(roots []*PathInfo, verbose bool) {
	var wg sync.WaitGroup
	filesize := make(chan *PathInfo)
	start := time.Now()

	for _, root := range roots {
		wg.Add(1)
		go root.walkDir(root.Name, filesize, &wg)
	}

	go func() {
		wg.Wait()
		close(filesize)
	}()

	var ticker <-chan time.Time
	if verbose {
		ticker = time.Tick(300 * time.Millisecond)
	}

loop:
	for {
		select {
		case <-ticker:
			printDiskUsage(roots, start)
		case _, ok := <-filesize:
			if !ok {
				break loop
			}
		}
	}
	printDiskUsage(roots, start)
}

func printDiskUsage(roots []*PathInfo, start time.Time) {
	for _, root := range roots {
		fmt.Printf(
			"--%s files %d, size %.2f GB, time %dms--\n",
			root.Name,
			root.FileCount,
			float64(root.Size)/1e9,
			time.Since(start).Milliseconds())
	}
}
