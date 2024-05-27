// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/

package main

func main() {
	nums := []int{3, 5}
	println(specialArray(nums))
}

func specialArray(nums []int) int {
	f := map[int]int{}
	for _, num := range nums {
		_, found := f[num]
		if !found {
			f[num] = 0
		}
		f[num]++
	}

	for x := 0; x <= len(nums); x++ {
		sum := 0
		for num, count := range f {
			if num >= x {
				sum += count
			}
		}
		if sum == x {
			return x
		}
	}
	return -1
}
