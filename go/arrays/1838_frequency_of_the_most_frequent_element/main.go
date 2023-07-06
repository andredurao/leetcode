// https://leetcode.com/problems/frequency-of-the-most-frequent-element/description/
package main

import (
	"math"
	"sort"
)

func main() {

}

// My initial attempt to solve this was trying to abstract the k value
// as the max area of a right triangle, from nums = [3,4,2,6,1,5] and k = 6
// after sorting nums and plotting the values
// |     #
// |    ##
// |kkk###
// |kk####
// |k#####
// |######
// +123456
// k would be able to fill up the space necessary for 1, 2 and 3 and with that
// the highest frequency would be 4.Unfortunatelly the nums set wouldn't perform as
// a linear function and is not possible to solve using this analogy

func maxFrequency(nums []int, k int) int {
	sort.IntSlice.Sort(nums)
	l := 0
	total := 0
	curr := 0
	for r, target := range nums {
		curr += target
		for ((r-l+1)*target)-curr > k {
			curr -= nums[l]
			l++
		}
		total = int(math.Max(float64(total), float64(r-l+1)))
	}

	return total
}
