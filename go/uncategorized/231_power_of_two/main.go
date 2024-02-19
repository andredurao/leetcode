// https://leetcode.com/problems/power-of-two/description

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOfTwo(3))
}

// I've used the log2 math function but
// maybe using bit masks could be even faster because powers of 2
// start with a 1 and end with 0 to n zeroes: 1, 10, 100, 1000, ...
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	l2 := math.Log2(float64(n))
	exp2 := math.Exp2(math.Trunc(l2))
	return int(exp2) == n
}
