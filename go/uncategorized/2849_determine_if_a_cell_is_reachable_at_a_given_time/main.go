// https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/

package main

import (
	"fmt"
	"math"
)

func main() {
	result := isReachableAtTime(2, 4, 7, 7, 6)
	// result := isReachableAtTime(1, 2, 1, 2, 1)
	fmt.Println(result)
}

// My first plan was to fetch the linear coefficents of the equation and check the distance but that
// was more than what I needed, I figured it out later that the minimum distance travelled was just
// the maximum diff of x or y coords.
func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	distY := math.Abs(float64(fy - sy))
	distX := math.Abs(float64(fx - sx))
	minDiff := int(math.Max(distX, distY))

	// case start and finish are the same coordinates, but there's no time to travel back
	if minDiff == 0 && t == 1 {
		return false
	}

	return minDiff <= t
}
