//

package main

import "fmt"

func main() {
	s := "aaa"
	res := partition(s)
	fmt.Println(res)
}

func partition(s string) [][]string {
	res := [][]string{}

	for i := range s {
		fmt.Println("size", i+1)
		res = append(res, []string{})
		doPartition(&s, &res, i, i+1, 0)
	}
	// fmt.Println(isPalindrome(&s, 0, 0))

	return res
}

func doPartition(s *string, res *[][]string, pos int, size int, start int) {
	if size == 0 {
		return
	}
	for l := start; l <= len(*s)-size; l++ {
		if isPalindrome(s, l, l+size-1) {
			(*res)[pos] = append((*res)[pos], (*s)[l:(l+size)])
		} else {
			doPartition(s, res, pos, size-1, l)
		}
	}
}

func isPalindrome(s *string, l int, r int) bool {
	res := true

	for l < r {
		if (*s)[l] != (*s)[r] {
			return false
		}
		l++
		r--
	}

	return res
}
