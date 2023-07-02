package bank

import "sync"

var (
	balance int
	mu      sync.Mutex
)

func Deposit(amount int) {
	defer mu.Unlock()
	mu.Lock()
	balance += amount

}

func deposit(amount int) {
	balance += amount
}

func Balance() int {
	defer mu.Unlock()
	mu.Lock()
	return balance
}

func WithDraw(amount int) bool {
	defer mu.Unlock()
	mu.Lock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
