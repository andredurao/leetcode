// https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
package main

import (
	"fmt"
	"strings"
)

func main() {

	garbage := []string{"G", "P", "GP", "GG"}
	travel := []int{2, 4, 3}
	result := garbageCollection(garbage, travel)
	fmt.Println(result)
}

func garbageCollection(garbage []string, travel []int) int {
	types := []rune{'P', 'M', 'G'}
	lastPlaces := map[rune]int{}

	// map last places
	for i, house := range garbage {
		for _, garbageType := range types {
			if strings.Contains(house, string(garbageType)) {
				lastPlaces[garbageType] = i
			}
		}
	}

	total := 0
	for _, garbageType := range types {
		_, found := lastPlaces[garbageType]
		if !found {
			continue
		}
		for i := 0; i <= lastPlaces[garbageType]; i++ {
			if i > 0 {
				total += travel[i-1]
			}
			for _, char := range garbage[i] {
				if char == garbageType {
					total++
				}
			}
		}
	}

	return total
}
