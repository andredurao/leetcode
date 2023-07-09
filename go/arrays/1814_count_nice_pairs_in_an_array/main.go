package main

import "fmt"

func main() {
	nums := []int{42, 11, 1, 97}
	// nums := []int{13, 10, 35, 24, 76}
	fmt.Println(countNicePairs(nums))
}

func countNicePairs(nums []int) int {
	revs := make([]int, len(nums))
	for i, n := range nums {
		revs[i] = n - rev(n)
		// fmt.Println(n, revs[i])
	}

	total := 0

	counts := map[int]int{}

	for _, n := range revs {
		_, found := counts[n]

		if !found {
			counts[n] = 0
		}
		total += counts[n]
		counts[n]++
	}

	return total % (int(1e9) + 7)
}

func rev(n int) int {
	val := 0
	for n > 0 {
		val = val*10 + n%10
		n /= 10
	}
	return val
}
