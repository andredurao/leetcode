//

package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 4}
	res := sortArray(nums)
	fmt.Println(res)
}

// --------------------------------

func sortArray(nums []int) []int {
	for i := (len(nums)/2 - 1); i >= 0; i-- {
		heapify(&nums, i, len(nums))
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(&nums, 0, i)
	}
	return nums
}

func heapify(nums *[]int, i int, size int) {
	maxIndex := i
	l, r := 2*i+1, 2*i+2

	if l < size && (*nums)[maxIndex] < (*nums)[l] {
		maxIndex = l
	}
	if r < size && (*nums)[maxIndex] < (*nums)[r] {
		maxIndex = r
	}
	if maxIndex != i {
		(*nums)[i], (*nums)[maxIndex] = (*nums)[maxIndex], (*nums)[i]
		heapify(nums, maxIndex, size)
	}
}
