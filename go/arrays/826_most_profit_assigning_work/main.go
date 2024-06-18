// https://leetcode.com/problems/most-profit-assigning-work

package main

import (
	"sort"
)

func main() {
	// difficulty := []int{2, 4, 6, 8, 10}
	// profit := []int{10, 20, 30, 40, 50}
	// worker := []int{4, 5, 6, 7}
	// difficulty := []int{85, 47, 57}
	// profit := []int{24, 66, 99}
	// worker := []int{40, 25, 25}
	difficulty := []int{13, 37, 58}
	profit := []int{4, 90, 96}
	worker := []int{34, 73, 45}
	res := maxProfitAssignment(difficulty, profit, worker)
	println(res)
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	total := 0

	// zip difficulty and profit into a single sorted [][]int
	jobs := [][]int{}
	for i := range difficulty {
		jobs = append(jobs, []int{difficulty[i], profit[i]})
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][0] < jobs[j][0] })
	// update the profits to their previous max values, if a worker is capable of difficulty = 2,
	// it should be able to deal with difficulty = 1 also
	maxProfit := jobs[0][1]
	for i := range jobs {
		maxProfit = max(maxProfit, jobs[i][1])
		jobs[i][1] = maxProfit
	}

	for _, ability := range worker {
		// check the highest profit / ability-wise for each worker using binary search rightmost insert point
		curProfit, l, r := 0, 0, len(jobs)-1
		for l <= r {
			mid := (l + r) / 2
			if jobs[mid][0] <= ability {
				curProfit = max(curProfit, jobs[mid][1])
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		total += curProfit
	}

	return total
}
