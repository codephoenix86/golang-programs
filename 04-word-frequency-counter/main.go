package main

import "fmt"

func GetWordFrequencies(words []string) map[string]int {
	fmap := make(map[string]int)

	for _, word := range words {
		fmap[word]++
	}

	return fmap
}

func main() {
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	fmap := GetWordFrequencies(words)

	for word, frequency := range fmap {
		fmt.Printf("%s: %d\n", word, frequency)
	}
}
