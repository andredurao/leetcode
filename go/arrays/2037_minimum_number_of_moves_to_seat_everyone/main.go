// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone

package main

import "sort"

func main() {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	println(minMovesToSeat(seats, students))
}

func minMovesToSeat(seats []int, students []int) int {
	sort.IntSlice.Sort(seats)
	sort.IntSlice.Sort(students)
	total := 0

	for i := range seats {
		total += abs(students[i] - seats[i])
	}

	return total
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
