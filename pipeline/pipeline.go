package pipeline

import "fmt"

func Pipe() {
	naturals := make(chan int)
	squares := make(chan int)

	go accumulate(naturals)
	go squarer(squares, naturals)

	for x := range squares {
		fmt.Println(x)
	}
}

func accumulate(naturals chan<- int) {
	for i := 1000; i > 0; i-- {
		naturals <- i
	}
	close(naturals)
}

func squarer(squares chan<- int, natural <-chan int) {
	for x := range natural {
		squares <- x * x
	}
	close(squares)
}
