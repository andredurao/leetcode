//

package main

import "fmt"

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	totalProduct := 1
	totalProductWithoutZero := 1
	zeroCount := 0
	for _, num := range nums {
		totalProduct *= num
		if num == 0 {
			zeroCount++
		} else {
			totalProductWithoutZero *= num
		}
	}
	if zeroCount > 1 {
		totalProductWithoutZero = 0
	}
	res := make([]int, len(nums))
	for i, num := range nums {
		if num == 0 {
			res[i] = totalProductWithoutZero
		} else {
			res[i] = totalProduct / num
		}
	}
	return res
}
