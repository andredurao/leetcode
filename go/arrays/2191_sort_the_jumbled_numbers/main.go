//

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	mapping := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
	nums := []int{991, 338, 38}

	// mapping := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// nums := []int{789, 456, 123}

	// mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := sortJumbled(mapping, nums)
	fmt.Println(res)
}

// -----------------------------------------------

func sortJumbled(mapping []int, nums []int) []int {
	zip := make([][]int, len(nums))

	for i, num := range nums {
		zip[i] = []int{num, convert(&mapping, num)}
	}

	sort.SliceStable(zip, func(i, j int) bool { return zip[i][1] < zip[j][1] })
	for i := range zip {
		nums[i] = zip[i][0]
	}

	return nums
}

func convert(mapping *[]int, num int) int {
	str := ""
	for _, ch := range strconv.Itoa(num) {
		str += fmt.Sprintf("%d", (*&*mapping)[ch-'0'])
	}
	res, _ := strconv.Atoi(str)
	return res
}
