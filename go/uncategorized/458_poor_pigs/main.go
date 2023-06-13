// https://leetcode.com/problems/poor-pigs/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	result := poorPigs(4, 15, 15)
	fmt.Println(result)
}

// from editorial
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	states := (minutesToTest / minutesToDie) + 1

	x := math.Log(float64(buckets)) / math.Log(float64(states))

	return int(math.Ceil(x - float64(1e-10)))
}
