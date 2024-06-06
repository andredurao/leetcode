// https://leetcode.com/problems/hand-of-straights/

package main

import "sort"

func main() {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	groupSize := 3
	println(isNStraightHand(hand, groupSize))
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	f := map[int]int{}
	for _, card := range hand {
		f[card]++
	}

	cards := make([]int, len(f))
	i := 0
	for k := range f {
		cards[i] = k
		i++
	}

	sort.IntSlice.Sort(cards)
	card := cards[0]
	group := 0
	sloth := 0
	for i := 0; i < len(hand); i++ {
		sloth = i % groupSize
		group = i / groupSize
		if sloth == 0 && group > 0 {
			// search smaller card for new hands that just started
			for _, c := range cards {
				if count, _ := f[c]; count > 0 {
					card = c
					break
				}
			}
		}

		if val, found := f[card]; !found || val == 0 {
			return false
		}
		f[card]--
		card++
	}

	return true
}
