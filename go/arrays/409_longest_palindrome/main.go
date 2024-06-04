// https://leetcode.com/problems/longest-palindrome/

package main

func main() {
	s := "abccccdd"
	println(longestPalindrome(s))
}

func longestPalindrome(s string) int {
	f := map[rune]int{}

	for _, ch := range s {
		if _, found := f[ch]; !found {
			f[ch] = 0
		}
		f[ch]++
	}

	total := len(s)

	odds := 0
	for _, count := range f {
		if count%2 == 1 {
			odds++
		}
	}

	for odds > 1 {
		total--
		odds--
	}

	return total
}
