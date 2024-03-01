// https://leetcode.com/problems/maximum-odd-binary-number/description/

package main

import "fmt"

func main() {
	fmt.Println(maximumOddBinaryNumber("010"))
}

func maximumOddBinaryNumber(s string) string {
	digits := make([]rune, len(s))
	l := 0
	r := len(s) - 2

	for _, c := range s {
		// fmt.Println(c, l, r, digits)
		if c == '1' {
			if digits[len(s)-1] == 0 {
				digits[len(s)-1] = c
			} else {
				digits[l] = c
				l++
			}
		} else {
			digits[r] = c
			r--
		}
	}

	return string(digits)
}
