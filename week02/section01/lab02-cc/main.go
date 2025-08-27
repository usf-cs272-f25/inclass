package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

// normalizeWord converts a word to lowercase and removes all punctuation,
// keeping only letters and digits
func normalizeWord(word string) string {
	var result strings.Builder
	for _, r := range word {
		// Keep only alphanumeric characters
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

func main() {
	// Read the entire text file
	content, err := ioutil.ReadFile("chapter1.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert bytes to string
	text := string(content)
	
	// Split text into words using whitespace as delimiter
	words := strings.Fields(text)
	
	// Create a map to store word frequencies
	wordFreq := make(map[string]int)
	
	// Process each word: normalize and count frequencies
	for _, word := range words {
		normalized := normalizeWord(word)
		// Only count non-empty normalized words
		if normalized != "" {
			wordFreq[normalized]++
		}
	}
	
	// Find the word with maximum frequency
	var maxWord string
	maxCount := 0
	
	for word, count := range wordFreq {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}
	
	// Output the most frequent word and its count
	fmt.Printf("%s %d\n", maxWord, maxCount)
}