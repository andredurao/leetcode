// https://leetcode.com/problems/diagonal-traverse-ii/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	nums := [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}
	// nums := [][]int{{14, 12, 19, 16, 9}, {13, 14, 15, 8, 11}, {11, 13, 1}}
	res := findDiagonalOrder(nums)
	fmt.Printf("%v\n", nums)
	fmt.Printf("%v\n", res)
}

func findDiagonalOrder(nums [][]int) []int {
	groups := map[int][][]int{}
	for i, row := range nums {
		for j, num := range row {
			sum := i + j
			_, found := groups[sum]
			if !found {
				groups[sum] = [][]int{}
			}
			// groups[sum] = append(groups[sum], []int{i, num})
			groups[sum] = append([][]int{{i, num}}, groups[sum]...)
		}
	}
	// fmt.Printf("%v\n", groups)
	keys := make([]int, len(groups))
	i := 0
	for k := range groups {
		keys[i] = k
		i++
	}
	sort.IntSlice.Sort(keys)
	result := []int{}
	for _, key := range keys {
		for _, group := range groups[key] {
			// fmt.Println("k", key, "gr", group)
			result = append(result, group[1])
		}
	}
	return result
}
