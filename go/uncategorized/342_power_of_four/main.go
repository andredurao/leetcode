package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	result := isPowerOfFour(n)
	fmt.Println(result)
}

// An integer n is a power of four, if there exists an integer x such that n == 4^x.
// log4 x = n ;
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	logResult := math.Log(float64(n)) / math.Log(4.0)

	return logResult == float64(int(logResult))
}
