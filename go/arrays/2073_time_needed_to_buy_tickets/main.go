// https://leetcode.com/problems/time-needed-to-buy-tickets/description/

package main

import "fmt"

func main() {
	tickets := []int{2, 3, 2}
	k := 2
	res := timeRequiredToBuy(tickets, k)
	fmt.Println(res)
}

func timeRequiredToBuy(tickets []int, k int) int {
	total := 0

	for {
		for i := range tickets {
			if tickets[i] > 0 {
				tickets[i]--
				total++
			}
			if i == k && tickets[i] == 0 {
				return total
			}
		}
	}
}
