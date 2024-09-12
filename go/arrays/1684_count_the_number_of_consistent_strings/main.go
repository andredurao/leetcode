//

package main

import "fmt"

func main() {
	// allowed := "ab"
	// words := []string{"ad", "bd", "aaab", "baa", "badab"}

	allowed := "abc"
	words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}

	// allowed := "cad"
	// words := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}

	res := countConsistentStrings(allowed, words)
	fmt.Println(res)
}

// ------------------------------

func countConsistentStrings(allowed string, words []string) int {
	allowedMap := make([]bool, 26)
	for _, ch := range allowed {
		allowedMap[int(ch-'a')] = true
	}

	total := 0

	for _, word := range words {
		if checkWord(&word, &allowedMap) {
			total++
		}
	}

	return total
}

func checkWord(word *string, allowedMap *[]bool) bool {
	for _, ch := range *word {
		if !(*allowedMap)[int(ch-'a')] {
			return false
		}
	}
	return true
}
