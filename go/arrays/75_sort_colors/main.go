// https://leetcode.com/problems/sort-colors/

package main

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}

func sortColors(nums []int) {
	f := []int{0, 0, 0}
	for _, num := range nums {
		f[num]++
	}

	i := 0
	for num := 0; num < 3; num++ {
		for j := 0; j < f[num]; j++ {
			nums[i] = num
			i++
		}
	}
	// fmt.Println(nums)
}
