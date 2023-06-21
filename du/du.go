package du

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, filesize chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, filesize, wg)
		} else {
			stat, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "can't get info about %s %v\n", entry.Name(), err)
			}
			filesize <- stat.Size()
		}
	}
}

func dirents(dir string) []os.DirEntry {
	//sema := make(chan struct{}, 5)
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read %s %v\n", dir, err)
		return nil
	}
	return entries
}

func ScanDirs(roots []string, verbose bool) {
	var wg sync.WaitGroup
	filesize := make(chan int64)
	start := time.Now()

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, filesize, &wg)
	}

	go func() {
		wg.Wait()
		close(filesize)
	}()

	var ticker <-chan time.Time
	var nfiles, nbytes int64
	if verbose {
		ticker = time.Tick(300 * time.Millisecond)
	}

loop:
	for {
		select {
		case <-ticker:
			printDiskUsage(nfiles, nbytes, start)
		case size, ok := <-filesize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}
	printDiskUsage(nfiles, nbytes, start)
}

func printDiskUsage(nfiles, nbytes int64, start time.Time) {
	fmt.Printf(
		"files %d, size %.2f GB, time %dms \n",
		nfiles,
		float64(nbytes)/1e9,
		time.Since(start).Milliseconds())
}
