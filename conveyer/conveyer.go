package conveyer

import (
	"log"
	"sync"
)

func Gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, i := range nums {
		out <- i
	}
	close(out)
	return out
}

func Square(done <-chan struct{}, nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range nums {
			select {
			case out <- i * i:
			case <-done:
				log.Print("canceling")
				return
			}
		}
	}()
	return out
}

func Merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range cs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for i := range ch {
				select {
				case out <- i:
				case <-done:
					log.Print("canceling")
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
