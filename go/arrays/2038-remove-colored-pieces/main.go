// https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/description/

package main

import (
	"fmt"
)

func main() {
	colors := "AAABABB"
	result := winnerOfGame(colors)
	fmt.Printf("%v\n", result)
}

func winnerOfGame(colors string) bool {
	alice := 0
	bob := 0

	for i, _ := range colors {
		if i == 0 || i == len(colors)-1 {
			continue
		}
		if colors[i-1] == colors[i] && colors[i] == colors[i+1] {
			if colors[i] == 'A' {
				alice++
			} else {
				bob++
			}
		}

	}
	return alice-bob >= 1
}
