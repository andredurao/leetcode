// https://leetcode.com/problems/custom-sort-string/description/

package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	order := "cba"
	s := "abcd"

	fmt.Println(customSortString(order, s))
}

func customSortString(order string, s string) string {
	var buffer bytes.Buffer
	charMap := map[rune]int{}

	// split chars and their frequencies
	for _, char := range s {
		_, found := charMap[char]
		if !found {
			charMap[char] = 0
		}
		charMap[char]++
	}

	// rebuild string in custom order
	for _, orderChar := range order {
		qty, found := charMap[orderChar]
		if found {
			for i := 0; i < qty; i++ {
				buffer.WriteRune(orderChar)
			}
			delete(charMap, orderChar)
		}
	}

	// adds remaining chars
	remainingOrder := []int{}
	for key := range charMap {
		remainingOrder = append(remainingOrder, int(key))
	}
	sort.IntSlice.Sort(remainingOrder)

	for _, orderChar := range remainingOrder {
		qty, _ := charMap[rune(orderChar)]
		for i := 0; i < qty; i++ {
			buffer.WriteRune(rune(orderChar))
		}
	}

	return buffer.String()
}
