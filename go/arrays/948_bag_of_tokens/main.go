//

package main

import (
	"fmt"
	"sort"
)

func main() {
	tokens := []int{100, 200}
	power := 150
	fmt.Println(bagOfTokensScore(tokens, power))
}

// Face-up: If your current power is at least tokens[i], you may play token i, losing tokens[i] power and gaining 1 score.
// Face-down: If your current score is at least 1, you may play token i, gaining tokens[i] power and losing 1 score.

func bagOfTokensScore(tokens []int, power int) int {
	l := 0
	r := len(tokens) - 1
	score := 0

	sort.IntSlice.Sort(tokens)

	if len(tokens) == 0 {
		return 0
	}

	for l <= r {
		if power >= tokens[l] { // face up
			fmt.Println("face up", tokens[l], power, l, r, score)
			power -= tokens[l]
			score++
			l++
		} else if (r - l) > 1 { // face down
			fmt.Println("face down", tokens[r], power, l, r, score)
			if score == 0 {
				return 0
			}
			power += tokens[r]
			score--
			r--
		} else {
			r--
		}
	}

	return score
}
