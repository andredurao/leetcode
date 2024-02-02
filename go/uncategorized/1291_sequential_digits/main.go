// https://leetcode.com/problems/sequential-digits/?envType=daily-question&envId=2024-02-02
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	res := sequentialDigits(1000, 13000)
	fmt.Println(res)
	seqs := sequences(1)
	fmt.Println(seqs)
}

func sequentialDigits(low int, high int) []int {
	lowDigits := len(strconv.Itoa(low))
	highDigits := len(strconv.Itoa(high))

	res := []int{}
	for size := lowDigits; size <= highDigits; size++ {
		for val := range sequences(size) {
			if val <= high && val >= low {
				res = append(res, val)
			}
		}
	}

	sort.IntSlice.Sort(res)

	return res
}

func sequences(size int) map[int]struct{} {
	res := map[int]struct{}{}

	if size == 1 {

	}

	init := 0
	for i := 1; i <= size; i++ {
		init *= 10
		init += i
	}

	res[init] = struct{}{}
	// start at 1 because the initial value (123...) has been set already
	// 9 - size => the amount of valid sequences decreases according to the size
	for i := 1; i <= (9 - size); i++ {
		val := nextSequence(init, size)
		res[val] = struct{}{}
		init = val
	}

	return res
}

func nextSequence(val int, size int) int {
	if size == 1 {
		val++
		return val
	}

	str := strconv.Itoa(val)
	str = str[1:]
	res, _ := strconv.Atoi(str)
	lastDigit, _ := strconv.Atoi(string(str[len(str)-1]))
	res *= 10
	res += lastDigit
	res++

	return res
}
