// https://leetcode.com/problems/find-unique-binary-string/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []string{"01", "10"}
	fmt.Println(findDifferentBinaryString(nums))
}

func findDifferentBinaryString(nums []string) string {
	n := len(nums[0])
	maxValue := (1 << n) - 1

	valuesMap := map[int]string{}
	for _, num := range nums {
		value, _ := strconv.ParseInt(num, 2, 64)
		valuesMap[int(value)] = num
	}

	result := ""
	for i := 0; i <= maxValue; i++ {
		_, found := valuesMap[i]
		if !found {
			result = fmt.Sprintf("%b", i)
			break
		}
	}

	for len(result) < n {
		result = "0" + result
	}

	return result
}
