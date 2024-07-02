// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	f1 := map[int]int{}
	f2 := map[int]int{}

	for _, num := range nums1 {
		f1[num]++
	}

	for _, num := range nums2 {
		f2[num]++
	}
	res := []int{}

	// values in nums1 that are present in nums2
	for num, count := range f1 {
		if f2[num] > 0 {
			for i := 0; i < min(count, f2[num]); i++ {
				res = append(res, num)
			}
			delete(f2, num)
		}
		delete(f1, num)
	}

	// values in nums2 that are present in nums1
	for num, count := range f2 {
		if f1[num] > 0 {
			for i := 0; i < min(count, f1[num]); i++ {
				res = append(res, num)
			}
			delete(f1, num)
		}
		delete(f2, num)
	}

	return res
}
