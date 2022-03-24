package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text []string) map[string]int {
	linecount := 35
	threads := int(math.Max(0, float64(len(text)/linecount)) + 1) // + 1 to not miss last partial line bunch

	var wg sync.WaitGroup
	c := make(chan map[string]int, threads)

	go func() {
		for t := 0; t < threads; t++ {
			wg.Add(1)

			go func(min int, max int) {
				defer wg.Done()
				freqs := make(map[string]int)
				for i := min; i < max; i++ {
					for _, word := range strings.Fields(text[i]) {
						freqs[strings.Trim(strings.Trim(strings.ToLower(word), "."), ",")]++
					}
				}
				c <- freqs // Return map to main routine
			}(t*linecount, int(math.Min(float64(len(text)), float64((t+1)*linecount))))
		}
		wg.Wait()
		defer close(c) // Close channel after all subthreads are done
	}()

	wg.Wait()

	// Merge together all counts
	merged := make(map[string]int)
	for m := range c {
		for k, v := range m {
			merged[k] += v
		}
	}

	return merged
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text []string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read file line by line to more easy parallelizisece
	file, err := os.Open("loremipsum.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var data []string

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			data = append(data, line)
		}
	}

	numRuns := 100
	runtimeMillis := benchmark(data, numRuns)
	printResults(runtimeMillis, numRuns)
}
