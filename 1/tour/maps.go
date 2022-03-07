package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	output := make(map[string]int)

	for _, word := range strings.Fields(s) {
		output[word]++
	}

	return output
}

func main() {
	wc.Test(WordCount)
}
