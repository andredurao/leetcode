package main

import (
	"fmt"
)

func main() {
	result := kthGrammar(5, 12)
	fmt.Println(result)
}

// inspired by the idea on editorial and discussion
func kthGrammar(n int, k int) int {
	return search(n, k)
}

func search(n int, k int) int {
	if n == 1 {
		return 0
	}

	count := 1 << (n - 1)
	half := count / 2
	if k > (count / 2) {
		return 1 - search(n, k-half) // invert the bit on response and shift left to the 1st half
	}
	return search(n-1, k)
}
