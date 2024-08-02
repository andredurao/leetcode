//

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0, 1, 0, 1, 1, 0, 0}
	res := minSwaps(nums)
	fmt.Println(res)
}

// ------------------------------------------

func minSwaps(nums []int) int {
	nnums := append(nums, nums...)
	ones := 0
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}

	zeroes := 0
	for i := 0; i < ones; i++ {
		if nnums[i] == 0 {
			zeroes++
		}
	}

	res := math.MaxInt32

	for i := 0; i <= len(nnums)-ones; i++ {
		if i > 0 {
			if nnums[i-1] == 0 {
				zeroes--
			}
			if nnums[i+ones-1] == 0 {
				zeroes++
			}
		}
		res = min(res, zeroes)
	}

	return res
}
