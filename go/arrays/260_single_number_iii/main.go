// https://leetcode.com/problems/single-number-iii/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) []int {
	f := map[int]int{}
	res := []int{}

	for _, num := range nums {
		if _, found := f[num]; !found {
			f[num] = 0
		}
		f[num]++
	}

	for num, count := range f {
		if count == 1 {
			res = append(res, num)
		}
	}

	return res
}
