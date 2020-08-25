package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq("wordfreq.txt")
}

func wordfreq(file string) {
	wordCounts := make(map[string]int)
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		currentWord := scanner.Text()
		_, found := wordCounts[currentWord]

		if found {
			wordCounts[currentWord]++
		} else {
			wordCounts[currentWord] = 1
		}
	}

	for word, count := range wordCounts {
		fmt.Println(word, ": ", count)
	}
}
