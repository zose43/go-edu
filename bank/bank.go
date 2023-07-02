package bank

var (
	deposit = make(chan int)
	balance = make(chan int)
	payoff  = make(chan *withdraw)
)

type withdraw struct {
	amount int
	ch     chan<- bool
}

func Deposit(amount int) {
	deposit <- amount
}

func Balance() int {
	return <-balance
}

func WithDraw(amount int) bool {
	ch := make(chan bool)
	wd := withdraw{ch: ch, amount: amount}
	payoff <- &wd
	return <-ch
}

func teller() {
	var balances int
	for {
		select {
		case amount := <-deposit:
			balances += amount
		case wd := <-payoff:
			if balances >= wd.amount {
				balances -= wd.amount
				wd.ch <- true
			} else {
				wd.ch <- false
			}
		case balance <- balances:
		}
	}
}

func init() {
	go teller()
}
