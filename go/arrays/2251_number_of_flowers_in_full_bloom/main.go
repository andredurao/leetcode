// https://leetcode.com/problems/number-of-flowers-in-full-bloom/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// flowers := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	// people := []int{2, 3, 7, 11}
	flowers := [][]int{{21, 34}, {17, 37}, {23, 43}, {17, 46}, {37, 41}, {44, 45}, {32, 45}}
	people := []int{31, 41, 10, 12}
	result := fullBloomFlowers(flowers, people)
	fmt.Printf("%v\n", result)
}

// From editorial solution
func fullBloomFlowers(flowers [][]int, people []int) []int {
	flowersInit := make([]int, len(flowers))
	flowersEnds := make([]int, len(flowers))
	for i, flower := range flowers {
		flowersInit[i] = flower[0]
		flowersEnds[i] = flower[1]
	}
	// sort by their initial times
	sort.IntSlice.Sort(flowersInit)
	sort.IntSlice.Sort(flowersEnds)

	fmt.Printf("%v\n", flowers)
	fmt.Printf("%v\n", flowersInit)
	fmt.Printf("%v\n", flowersEnds)
	fmt.Printf("%v\n", people)

	result := make([]int, len(people))

	for i, day := range people {
		// rightmost binary search
		initCount := sort.Search(len(flowersInit), func(i int) bool { return flowersInit[i] > day })
		// leftmost binary search
		endsCount := sort.Search(len(flowersEnds), func(i int) bool { return flowersEnds[i] >= day })
		result[i] = initCount - endsCount
	}
	return result
}
