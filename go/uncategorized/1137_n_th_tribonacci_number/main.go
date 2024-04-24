// https://leetcode.com/problems/n-th-tribonacci-number/description/

package main

func main() {
	println(tribonacci(25))
}

func tribonacci(n int) int {
	slots := make([]int, 4)
	slots[0] = 0
	slots[1] = 1
	slots[2] = 1

	if n < 3 {
		return slots[n]
	}

	for i := 0; i < n-2; i++ {
		slots[3] = slots[0] + slots[1] + slots[2]
		slots[0], slots[1], slots[2] = slots[1], slots[2], slots[3]
	}
	return slots[3]
}
