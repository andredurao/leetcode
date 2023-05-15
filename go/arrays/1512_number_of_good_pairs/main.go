// https://leetcode.com/problems/number-of-good-pairs/description/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	result := numIdenticalPairs(nums)
	fmt.Println(result)
}

func numIdenticalPairs(nums []int) int {
	first_occurrences := make(map[int]int)

	total := 0
	for _, v := range nums {
		fmt.Printf("%d %d\n", i, v)
		count, ok := first_occurrences[v]
		if ok {
			fmt.Println("found: ", v)
			total = total + count
			first_occurrences[v] = count + 1
		} else {
			fmt.Println("store: ", v)
			first_occurrences[v] = 1
		}
	}

	return total
}
