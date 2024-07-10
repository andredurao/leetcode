// https://leetcode.com/problems/crawler-log-folder/

package main

import (
	"fmt"
)

func main() {
	// logs := []string{"d1/", "d2/", "../", "d21/", "./"}
	// logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	// logs := []string{"d1/", "../", "../", "../"}
	logs := []string{"./", "../", "./"}

	res := minOperations(logs)
	fmt.Println(res)
}

func minOperations(logs []string) int {
	res := 0
	for _, path := range logs {
		dirName := path[:len(path)-1]
		if dirName == ".." {
			if res > 0 {
				res--
			}
		} else if dirName != "." {
			res++
		}
	}

	return res
}
