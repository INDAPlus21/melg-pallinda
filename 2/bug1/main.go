package main

import "fmt"

// The issue was that a unbuffered channel was used instead of a buffered one so it
// does not have capacity to store the value but rather exspects the value to be read
// instantly in another routine
func main() {
	ch := make(chan string, 1)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
