package spinner

import (
	"fmt"
	"time"
)

func Spinner(duration time.Duration) {
	for {
		for _, symb := range `-\|/` {
			fmt.Printf("\r%c", symb)
			time.Sleep(duration)
		}
	}
}
