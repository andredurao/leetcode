// https://leetcode.com/problems/isomorphic-strings/

package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	t := os.Args[2]
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	sCharmap := make([]rune, 256)
	tCharmap := make([]rune, 256)

	for i := range s {
		if sCharmap[rune(s[i])] == 0 && tCharmap[rune(t[i])] == 0 {
			sCharmap[rune(s[i])] = rune(t[i])
			tCharmap[rune(t[i])] = rune(s[i])
		} else {
			if sCharmap[rune(s[i])] != rune(t[i]) {
				return false
			}
		}
	}

	return true
}
