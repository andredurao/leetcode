// https://leetcode.com/problems/reveal-cards-in-increasing-order/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// input := []int{1, 6, 2, 5, 3, 7, 4}
	input := []int{17, 13, 11, 2, 3, 5, 7}
	output := deckRevealedIncreasing(input)
	fmt.Println(output)
	// input := []int{1, 9, 2, 6, 3, 8, 4, 7, 5}
	// input := []int{1, 5, 2, 7, 3, 6, 4, 8}
	// input := []int{1, 9, 2, 7, 3, 11, 4, 8, 5, 10, 6}
	final := unpack(output)
	fmt.Println(final)
}

func deckRevealedIncreasing(deck []int) []int {
	sort.IntSlice.Sort(deck)

	// initialize
	deckMap := make([]int, len(deck))
	for i := range deck {
		deckMap[i] = i
	}

	res := make([]int, len(deck))
	for _, val := range deck {
		res[deckMap[0]] = val
		deckMap = deckMap[1:]
		if len(deckMap) > 0 {
			head := deckMap[0]
			deckMap = deckMap[1:]
			deckMap = append(deckMap, head)
		}
	}

	return res
}

func unpack(input []int) []int {
	res := []int{}
	for len(input) > 0 {
		res = append(res, input[0])
		input = input[1:]
		if len(input) > 1 {
			head := input[0]
			input = input[1:]
			input = append(input, head)
		}
	}
	return res
}
