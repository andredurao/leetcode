//

package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := minBitFlips(3, 200)
	fmt.Println(res)
}

//  =======================================

func minBitFlips(start int, goal int) int {
	binValues := []string{strconv.FormatInt(int64(start), 2), strconv.FormatInt(int64(goal), 2)}
	total := 0
	end := max(len(binValues[0]), len(binValues[1]))

	sVal, gVal := '0', '0'
	for i := 0; i < end; i++ {
		if i >= len(binValues[0]) {
			sVal = '0'
		} else {
			sVal = rune(binValues[0][len(binValues[0])-i-1])
		}
		if i >= len(binValues[1]) {
			gVal = '0'
		} else {
			gVal = rune(binValues[1][len(binValues[1])-i-1])
		}
		// fmt.Println("i", i, "s", string(sVal), "g", string(gVal))
		if sVal != gVal {
			total++
		}
	}

	return total
}
