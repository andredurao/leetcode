// https://leetcode.com/problems/lemonade-change/

package main

import "fmt"

func main() {
	// bills := []int{5, 5, 10, 10, 20}
	// bills := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	bills := []int{5, 5, 5, 5, 20, 20, 5, 5, 20, 5}
	res := lemonadeChange(bills)
	fmt.Println(res)
}

// --------------------------------------

func lemonadeChange(bills []int) bool {
	cashier := []int{0, 0, 0}
	for _, bill := range bills {
		if bill == 5 {
			cashier[0]++
		} else if bill == 10 {
			cashier[1]++
			cashier[0]--
		} else if bill == 20 {
			if cashier[1] > 0 {
				cashier[2]++
				cashier[1]--
				cashier[0]--
			} else {
				cashier[2]++
				cashier[0] -= 3
			}
		}
		if cashier[0] < 0 || cashier[1] < 0 {
			return false
		}
	}

	return true
}
