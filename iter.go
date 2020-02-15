package iter

func R(n int) chan int {
	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}
