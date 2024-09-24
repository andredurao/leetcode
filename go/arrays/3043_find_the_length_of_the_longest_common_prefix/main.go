// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/

package main

import "fmt"

func main() {
	arr1 := []int{1, 10, 100}
	arr2 := []int{1000}
	res := longestCommonPrefix(arr1, arr2)
	fmt.Println(res)
}

// ================

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixesMap1 := map[int]struct{}{}
	for _, num := range arr1 {
		storePrefixes(num, prefixesMap1)
	}
	prefixesMap2 := map[int]struct{}{}
	for _, num := range arr2 {
		storePrefixes(num, prefixesMap2)
	}
	res := 0
	for prefix := range prefixesMap2 {
		if _, found := prefixesMap1[prefix]; found {
			res = max(res, prefixLen(prefix))
		}
	}
	return res
}

func storePrefixes(num int, prefixesMap map[int]struct{}) {
	for num > 0 {
		prefixesMap[num] = struct{}{}
		num /= 10
	}
}

func prefixLen(prefix int) (length int) {
	for prefix > 0 {
		length++
		prefix /= 10
	}
	return
}
