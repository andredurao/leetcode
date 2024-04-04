// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

package main

import "fmt"

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	fmt.Println(maxDepth(s))
}

func maxDepth(s string) int {
	total := 0
	depth := 0

	for _, ch := range s {
		if ch == '(' {
			depth++
		}
		if ch == ')' {
			depth--
		}
		total = max(total, depth)
	}

	return total
}
