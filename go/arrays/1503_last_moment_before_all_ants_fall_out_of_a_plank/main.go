package main

import (
	"fmt"
	"sort"
)

// TODO: Needs review, because is using too much memory

func main() {

	n := 4
	left := []int{4, 3}
	right := []int{0, 1}
	result := getLastMoment(n, left, right)
	fmt.Println(result)
}

func getLastMoment(n int, left []int, right []int) (result int) {
	for len(left) > 0 || len(right) > 0 {
		result++
		walk(n, &left, &right)
	}
	result--
	return
}

func walk(n int, left *[]int, right *[]int) {
	consume(n, left, -1)
	consume(n, right, +1)
}

func consume(n int, array *[]int, dir int) {
	sort.IntSlice.Sort(*array)
	for i, val := range *array {
		(*array)[i] = val + dir // -1(L) or +1(R)
	}

	// instead of creating new arrays I should change the range of the existing slice
	result := make([]int, 0)
	for _, val := range *array {
		if val >= 0 && val <= n {
			result = append(result, val)
		}
	}
	*array = result
}
