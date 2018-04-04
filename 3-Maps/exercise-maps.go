package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordList := map[string]int{}
	words := strings.Fields(s)

	for _, w := range words {
		if _, ok := wordList[w]; !ok {
			wordList[w] = 1
		} else {
			wordList[w] = wordList[w] + 1
		}
	}

	return wordList
}

func main() {
	wc.Test(WordCount)
}
