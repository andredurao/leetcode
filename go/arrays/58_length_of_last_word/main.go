// https://leetcode.com/problems/length-of-last-word/

package main

import "fmt"

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	res := 0
	i := len(s) - 1

	for s[i] == ' ' {
		i--
	}

	for ; i >= 0; i-- {
		if s[i] == ' ' {
			return res
		}
		res++
	}
	return res
}
