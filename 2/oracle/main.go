// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go answerQuestions(questions, answers)
	go makeProphecies(answers)
	go printAnswers(answers)

	return questions
}

func answerQuestions(questions <-chan string, answers chan<- string) {
	for question := range questions {
		go func() {
			time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second) // Randomize thinking time (1-3 seconds)
			prophecy(question, answers)
		}()
	}
}

func makeProphecies(answers chan<- string) {
	time.Sleep(time.Duration(rand.Intn(5)+10) * time.Second) // Every 5-15 seconds

	nonsense := []string{
		"If you close your eyes you will not see",
		"You future looks bright",
		"Every step you take is a step towards the end",
		"I can see that you have the right spirit inside of you",
		"Not even the gods are perfect",
		"When night becomes day, yesterday will be forever gone",
	}

	answers <- nonsense[rand.Intn(len(nonsense))]
	makeProphecies(answers) // Infinite loop
}

func printAnswers(ch <-chan string) {
	for s := range ch {
		fmt.Printf("%s: %s\n", star, s)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
