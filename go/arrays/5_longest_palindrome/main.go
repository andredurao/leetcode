package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	result := longestPalindrome(input)
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	for size := len(s); size > 0; size-- {
		for i := 0; i <= (len(s) - size); i++ {
			// fmt.Println(size, i, s[i:(i+size)])
			if isPalindrome(s[i:(i + size)]) {
				return s[i:(i + size)]
			}
		}
	}
	return string(s[0])
}

func isPalindrome(s string) bool {
	for i := 0; i < (len(s) / 2); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
