// https://leetcode.com/problems/first-unique-character-in-a-string/description/?envType=daily-question&envId=2024-02-05
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	charmap := map[rune][]int{}
	// fill a map of chars, their indexes and counts
	for index, ch := range s {
		_, found := charmap[ch]
		if !found {
			charmap[ch] = []int{index, 1}
		} else {
			charmap[ch][1]++
		}
	}

	indexes := []int{}
	for _, values := range charmap {
		if values[1] == 1 {
			indexes = append(indexes, values[0])
		}
	}

	if len(indexes) == 0 {
		return -1
	}

	sort.IntSlice.Sort(indexes)
	return indexes[0]
}
