// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/

package main

import "fmt"

func main() {
	steps := 3
	arrLen := 2
	result := numWays(steps, arrLen)
	fmt.Println(result)
}

func numWays(steps int, arrLen int) int {
	counter := 0
	dp(0, steps, arrLen, &counter)
	return counter
}

func dp(pos int, steps int, size int, counter *int) {
	if steps == 0 {
		if pos == 0 {
			*counter++
		}
		return
	}

	diff := pos - steps
	if diff == 0 {
		*counter++
		return
	} else if diff > 0 && pos != 0 {
		return
	}

	// left
	if pos > 0 {
		dp(pos-1, steps-1, size, counter)
	}

	// right
	if pos < (size - 1) {
		dp(pos+1, steps-1, size, counter)
	}

	// stay
	dp(pos, steps-1, size, counter)
}
