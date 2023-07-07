package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2, 3}
	fmt.Println(reductionOperations((nums)))
}

func reductionOperations(nums []int) int {
	frequencies := map[int]int{}

	for _, num := range nums {
		_, found := frequencies[num]
		if !found {
			frequencies[num] = 0
		}
		frequencies[num]++
	}

	total := 0
	keys := sortedKeys(frequencies)
	i := len(keys) - 1

	for i > 0 {
		max := keys[i]
		nextMax := keys[i-1]
		// fmt.Printf("bef %v\n", frequencies)
		total += frequencies[max]
		frequencies[nextMax] += frequencies[max]
		delete(frequencies, max)
		// fmt.Printf("aft %v\n\n", frequencies)
		i--
	}

	return total
}

// FIXME: Should have used map.Keys introduced in go 1.18 :(
// https://pkg.go.dev/golang.org/x/exp/maps#Keys
func sortedKeys(frequencies map[int]int) []int {
	res := make([]int, len(frequencies))
	i := 0
	for key, _ := range frequencies {
		res[i] = key
		i++
	}
	sort.IntSlice.Sort(res)
	return res
}
