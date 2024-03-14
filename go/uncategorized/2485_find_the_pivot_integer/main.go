// https://leetcode.com/problems/find-the-pivot-integer/

package main

func main() {
	println(pivotInteger(8))
}

func pivotInteger(n int) int {
	sigma := (n * (n + 1)) / 2
	l := 1
	r := n
	mid := 0
	lsum := 0
	rsum := 0

	for l <= r {
		mid = (l + r) / 2
		lsum = (mid * (mid + 1)) / 2
		rsum = sigma - ((mid-1)*mid)/2
		if lsum == rsum {
			return mid
		} else if lsum < rsum {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
