package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func normalizeWord(word string) string {
	var result strings.Builder
	result.Grow(len(word))
	
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

func processTextStream(reader io.Reader) (string, int) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	
	wordFreq := make(map[string]int)
	
	for scanner.Scan() {
		word := scanner.Text()
		normalized := normalizeWord(word)
		if normalized != "" {
			wordFreq[normalized]++
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading: %v\n", err)
		return "", 0
	}
	
	var mostFreqWord string
	var maxCount int
	
	for word, count := range wordFreq {
		if count > maxCount {
			mostFreqWord = word
			maxCount = count
		}
	}
	
	return mostFreqWord, maxCount
}

func main() {
	file, err := os.Open("call_of_the_wild.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	
	bufferedReader := bufio.NewReaderSize(file, 64*1024)
	
	mostFreqWord, count := processTextStream(bufferedReader)
	
	fmt.Printf("Most frequent word: %s\n", mostFreqWord)
	fmt.Printf("Count: %d\n", count)
}