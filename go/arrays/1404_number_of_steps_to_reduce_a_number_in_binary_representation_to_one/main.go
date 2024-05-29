// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/

package main

import (
	"os"
)

func main() {
	s := os.Args[1]
	println(numSteps(s))
}

func numSteps(s string) int {
	val := []byte(s)
	steps := 0

	for len(val) > 0 {
		// fmt.Println("S", string(val), len(val), steps)
		if len(val) == 1 && val[0] == '1' {
			break
		}
		end := len(val) - 1

		if val[end] == '0' {
			val = val[:end] // SHR
			steps++
		} else {
			j := end
			for j > 0 && val[j] == '1' {
				steps++
				j--
			}
			steps++
			if j == 0 {
				steps++
				val = []byte("1")
			} else {
				val[j] = '1'
				val = val[0 : j+1]
			}
		}
		// fmt.Println("E", string(val), len(val), steps, "\n")
	}

	return steps
}
