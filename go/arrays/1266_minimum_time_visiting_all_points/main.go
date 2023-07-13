// https://leetcode.com/problems/minimum-time-visiting-all-points/description/

package main

import "fmt"

func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	res := minTimeToVisitAllPoints(points)
	fmt.Println(res)
}

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	a := points[0]
	for i := 1; i < len(points); i++ {
		b := points[i]
		for !samePos(&a, &b) {
			if isSameAxis(&a, &b) {
				// move horizontally or vertically
				if a[0] == b[0] {
					a[1] += dir(&a, &b, 1)
				} else {
					a[0] += dir(&a, &b, 0)
				}
			} else {
				// move diagonally
				a[0] += dir(&a, &b, 0)
				a[1] += dir(&a, &b, 1)
			}
			time++
			// fmt.Println(a, b, time)
		}
	}

	return time
}

func isSameAxis(a, b *[]int) bool {
	return (*a)[0] == (*b)[0] || (*a)[1] == (*b)[1]
}

func samePos(a, b *[]int) bool {
	return (*a)[0] == (*b)[0] && (*a)[1] == (*b)[1]
}

func dir(a, b *[]int, axis int) int {
	if (*a)[axis] > (*b)[axis] {
		return -1
	}
	return 1
}
