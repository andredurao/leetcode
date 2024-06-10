// https://leetcode.com/problems/regular-expression-matching/

package main

import (
	"fmt"
)

func main() {
	s := "aab"
	p := "c*a*b"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	si, pi := 0, 0
	ch := ' '
	for si < len(s) && pi < len(p) {
		ch = rune(s[si])
		curP := rune(p[pi])
		if isChar(curP) {
			if ch != curP {
				return false
			}
			si++
			pi++
			continue
		}
		if curP == '.' {
			si++
			pi++
			continue
		}
		if curP == '*' {
			for si < len(s) && (p[pi-1] == s[si] || p[pi-1] == '.') {
				si++
			}
			pi++
			continue
		}
	}

	return si == len(s)
}

func isChar(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

func star(s *string, i int) bool {
	return false
}

func dot(s *string, i int) bool {
	return false
}
