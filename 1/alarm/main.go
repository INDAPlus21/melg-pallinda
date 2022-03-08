package main

import (
	"fmt"
	"sync"
	"time"
)

func Remind(text string, delay time.Duration) {
	fmt.Println("The time is " + time.Now().Format("15:04:05") + ": " + text)
	time.Sleep(delay)
	Remind(text, delay)
}

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go Remind("Time to eat", time.Second*10)
	go Remind("Time to work", time.Second*30)
	go Remind("Time to sleep", time.Second*60)

	wg.Wait() // Prevent main thread from sleeping
}
