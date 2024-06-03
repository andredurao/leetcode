// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/

package main

func main() {
	s := "coaching"
	t := "coding"
	println(appendCharacters(s, t))
}
func appendCharacters(s string, t string) int {
	si, ti := 0, 0

	// Use two variables to keep track of your location in s and t. If the characters match, increment both variables. Otherwise, only increment the variable for s.
	for si < len(s) {
		if ti < len(t) && s[si] == t[ti] {
			ti++
		}
		si++
	}

	// The remaining characters in t must be appended to the end of s.

	return len(t) - ti
}

// 1. Find the longest prefix of t that is a subsequence of s
func longestPrefix(s *string, t *string) (size int) {
	for i := range *s {
		if (*s)[i] == (*t)[size] {
			size++
		}
	}

	return
}
