package fibonachi

func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, y+x
	}
	return x
}

func FibByRec(n int) int {
	if n > 1 {
		return FibByRec(n-1) + FibByRec(n-2)
	}
	return n
}
