// https://leetcode.com/problems/sort-array-by-increasing-frequency/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 1, 2, 2, 2, 3}
	// nums := []int{2, 3, 1, 3, 2}
	nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}

	res := frequencySort(nums)
	fmt.Println(res)
}

func frequencySort(nums []int) []int {
	f := map[int]int{}
	for _, n := range nums {
		f[n]++
	}
	// zip nums and their frequencies
	zip := [][]int{}
	for k, v := range f {
		zip = append(zip, []int{k, v})
	}
	// sort as description
	sort.Slice(zip, func(i, j int) bool {
		if zip[i][1] == zip[j][1] {
			return zip[i][0] > zip[j][0]
		}
		return zip[i][1] < zip[j][1]
	})
	res := []int{}
	for _, vals := range zip {
		for i := 0; i < vals[1]; i++ {
			res = append(res, vals[0])
		}
	}
	return res
}
