// asd

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	f := map[int]int{}
	res := []int{}

	for _, x := range arr1 {
		f[x]++
	}

	for _, x := range arr2 {
		for i := 0; i < f[x]; i++ {
			res = append(res, x)
		}
		delete(f, x)
	}

	keys := []int{}
	for k := range f {
		keys = append(keys, k)
	}

	sort.IntSlice.Sort(keys)

	for _, x := range keys {
		for i := 0; i < f[x]; i++ {
			res = append(res, x)
		}
	}

	return res
}
