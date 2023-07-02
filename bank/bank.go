package bank

var (
	deposit = make(chan int, 10)
	balance = make(chan int, 10)
)

func Deposit(amount int) {
	deposit <- amount
}

func Balance() int {
	return <-balance
}

func teller() {
	var balances int
	for {
		select {
		case amount := <-deposit:
			balances += amount
		case balance <- balances:
		}
	}
}

func init() {
	go teller()
}
