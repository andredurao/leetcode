package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	size := 1 << len(nums)
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		vals := enabledBits(i)
		res[i] = make([]int, len(vals))
		for j := range res[i] {
			res[i][j] = nums[vals[j]]
		}
	}
	return res
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
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
