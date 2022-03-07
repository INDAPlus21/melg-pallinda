package main

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	out := 0
	for i := range a {
		out += a[i]
	}

	res <- out
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	return <-ch + <-ch
}
