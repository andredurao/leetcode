// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/

package main

import "fmt"

func main() {
	s := "ca"
	fmt.Println(minimumLength(s))
}

func minimumLength(s string) int {
	l := 0
	r := len(s) - 1
	for len(s) > 0 && s[l] == s[r] && l < r {
		ch := s[l]
		for l < (r-1) && ch == s[l+1] {
			l++
		}
		for r > (l+1) && ch == s[r-1] {
			r--
		}
		// cut off prefix and suffix

		s = s[(l + 1):r]
		l = 0
		r = len(s) - 1
	}

	return len(s)
}
