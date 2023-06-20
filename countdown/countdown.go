package countdown

import (
	"fmt"
	"os"
	"time"
)

func CountDown() {
	stop := make(chan struct{})
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		{
			_, _ = os.Stdin.Read(make([]byte, 1)) //caught enter
			stop <- struct{}{}
		}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
			// pending
		case <-stop:
			fmt.Println("Launch aborted")
			return
		}
	}

	ticker.Stop()
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
