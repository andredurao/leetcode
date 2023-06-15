package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	result := myAtoi(os.Args[1])
	fmt.Println(result)
}

func myAtoi(s string) int {
	var value int = 0
	positive := true
	inputing := false
	for _, c := range s {
		if !inputing && c == ' ' {
			continue
		}

		if !inputing && (c == '-' || c == '+') {
			inputing = true
			if c == '-' {
				positive = false
			}
			continue
		}

		if c >= '0' && c <= '9' {
			if !inputing {
				inputing = true
			}
			value = value * 10
			value = value + int(c-'0')
		}

		if inputing && (c < '0' || c > '9') {
			break
		}

		if !inputing {
			break
		}

		if (positive && value > math.MaxInt32) || (!positive && -value < math.MinInt32) {
			if positive {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}

	if !positive {
		value = -value
	}

	return value
}
