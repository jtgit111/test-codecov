package testCodecov

import (
	"regexp"
	"sort"
	"strings"
)

// WordFrequency holds a word and its frequency count.
type WordFrequency struct {
	Word  string
	Count int
}

// AnalyzeText takes a string input and returns a sorted list of word frequencies.
func AnalyzeText(input string) []WordFrequency {
	// Normalize the input by converting to lower case and removing punctuation
	re := regexp.MustCompile(`[^\w\s]`)
	normalizedInput := re.ReplaceAllString(strings.ToLower(input), "")

	// Split the input into words
	words := strings.Fields(normalizedInput)

	// Create a map to count word frequencies
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Convert the map to a slice of WordFrequency
	var frequencies []WordFrequency
	for word, count := range wordCount {
		frequencies = append(frequencies, WordFrequency{Word: word, Count: count})
	}

	// Sort the slice by frequency (descending)
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Count > frequencies[j].Count
	})

	return frequencies
}