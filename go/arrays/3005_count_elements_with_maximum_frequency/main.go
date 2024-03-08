// https://leetcode.com/problems/count-elements-with-maximum-frequency/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 1, 4}
	fmt.Println(maxFrequencyElements(nums))
}

func maxFrequencyElements(nums []int) int {
	frequencyMap := map[int]int{}
	maxFreq := 0

	for _, num := range nums {
		_, found := frequencyMap[num]
		if !found {
			frequencyMap[num] = 0
		}
		frequencyMap[num]++
		maxFreq = int(max(float64(maxFreq), float64(frequencyMap[num])))
	}

	res := 0
	for _, freq := range frequencyMap {
		if freq == maxFreq {
			res += freq
		}
	}

	return res
}
