// https://leetcode.com/problems/most-profit-assigning-work

package main

func main() {
	difficulty := []int{2, 4, 6, 8, 10}
	profit := []int{10, 20, 30, 40, 50}
	worker := []int{4, 5, 6, 7}
	res := maxProfitAssignment(difficulty, profit, worker)
	println(res)
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	total := 0

	for _, ability := range worker {
		// check the highest profit / ability-wise for each worker
		maxProfit := 0
		for i := range difficulty {
			if ability >= difficulty[i] {
				maxProfit = max(profit[i], maxProfit)
			}
		}
		total += maxProfit
	}

	return total
}
