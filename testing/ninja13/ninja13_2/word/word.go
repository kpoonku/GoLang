// Gets the words count for all the words and a map of word and its count
package word

import "strings"

// Gets the map of words and its word count
func getWordCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Get the count of the all the words in the given string
func count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
