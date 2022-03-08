package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// This program should go to 11, but it seemingly only prints 1 to 10.
// The issue was that the main thread shut down before the print thread
// was done so now it waits for both threads to complete. It also shuts
// down the channel when it's finished sending
func main() {
	ch := make(chan int, 11)
	wg.Add(1)
	go Print(ch)
	wg.Add(1)
	go func() {
		for i := 1; i <= 11; i++ {
			ch <- i
		}
		close(ch) // Shuts down after all messages have been recieved
		wg.Done()
	}()

	wg.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	defer wg.Done()
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}
