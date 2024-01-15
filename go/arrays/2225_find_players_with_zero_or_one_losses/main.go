package main

import (
	"fmt"
	"sort"
)

func main() {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	fmt.Println(findWinners(matches))
}

// Ported from a ruby submit that I did previously

func findWinners(matches [][]int) [][]int {
	winners := map[int]int{}
	loosers := map[int]int{}

	for _, match := range matches {
		winner := match[0]
		looser := match[1]

		winners[winner] = 1
		_, found := loosers[looser]
		if !found {
			loosers[looser] = 0
		}
		loosers[looser]++
	}

	oneTimeLoosers := []int{}
	for looser, qty := range loosers {
		delete(winners, looser)
		if qty == 1 {
			oneTimeLoosers = append(oneTimeLoosers, looser)
		}
	}

	winnersList := []int{}

	for winner := range winners {
		winnersList = append(winnersList, winner)
	}
	sort.IntSlice.Sort(winnersList)
	sort.IntSlice.Sort(oneTimeLoosers)

	res := [][]int{}
	res = append(res, winnersList)
	res = append(res, oneTimeLoosers)
	return res
}
