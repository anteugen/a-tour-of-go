package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

var word_count map[string]int

func WordCount(s string) map[string]int {
	words_map := strings.Fields(s)
	
	word_count = make(map[string]int)
	
	for _, word := range words_map {
		word_count[word]++
	}

	return word_count
}

func main() {
	wc.Test(WordCount)
}