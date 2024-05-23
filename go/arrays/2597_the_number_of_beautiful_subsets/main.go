package main

import (
	"fmt"
	"math"
)

func beautifulSubsets(nums []int, k int) int {
	total := 0
	size := 1 << len(nums)
	for i := 0; i < size; i++ {
		indexes := enabledBits(i)
		// fmt.Println("isB?", indexes)
		if isBeautiful(&nums, &indexes, k) {
			total++
		}
	}
	return total
}

func isBeautiful(nums *[]int, indexes *[]int, k int) bool {
	if len(*indexes) == 0 {
		return false
	}
	if len(*indexes) == 1 {
		return true
	}
	for ii := 0; ii < len(*indexes)-1; ii++ {
		for jj := ii + 1; jj < len(*indexes); jj++ {
			// fmt.Println("check", ii, jj, "/", (*nums)[ii], (*nums)[jj])
			if int(math.Abs(float64((*nums)[(*indexes)[ii]]-(*nums)[(*indexes)[jj]]))) == k {
				return false
			}
		}
	}
	return true
}

func enabledBits(c int) (res []int) {
	i := 0
	for c > 0 {
		if c%2 == 1 {
			res = append(res, i)
		}
		c >>= 1
		i++
	}

	return
}

func main() {
	nums := []int{2, 4, 6, 9}
	fmt.Println(beautifulSubsets(nums, 3))
}
