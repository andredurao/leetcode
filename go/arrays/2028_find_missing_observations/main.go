// https://leetcode.com/problems/find-missing-observations/

package main

import "fmt"

func main() {
	// res := missingRolls([]int{3, 2, 4, 3}, 4, 2)
	// res := missingRolls([]int{1, 5, 6}, 3, 4)
	// res := missingRolls([]int{1, 2, 3, 4}, 6, 4)
	// res := missingRolls([]int{6, 3, 4, 3, 5, 3}, 1, 6)
	// res := missingRolls([]int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}, 4, 40)
	res := missingRolls([]int{4, 2, 2, 5, 4, 5, 4, 5, 3, 3, 6, 1, 2, 4, 2, 1, 6, 5, 4, 2, 3, 4, 2, 3, 3, 5, 4, 1, 4, 4, 5, 3, 6, 1, 5, 2, 3, 3, 6, 1, 6, 4, 1, 3}, 2, 53)
	fmt.Println(res)
}

// ---------

func missingRolls(rolls []int, mean int, n int) []int {
	total := 0
	for _, roll := range rolls {
		total += roll
	}

	// 3+2+4+3 = 12    ; (12+x) / (4 + 2) = 4 ; 12+x = 4 * 6 ; 12+x = 24 ; x = 12
	//           total                        ; x = mean * (len(rolls) + n) - total
	diff := mean*(len(rolls)+n) - total
	fmt.Println(diff)
	res := []int{}
	if diff <= 0 || float32(diff)/float32(n) > 6.0 || float32(diff)/float32(n) < 1.0 {
		return res
	}
	avg := diff / n
	resRemaining := diff
	for i := 0; i < n; i++ {
		res = append(res, avg)
		resRemaining -= avg
	}
	for i := 0; i < n && resRemaining > 0; i++ {
		complement := 6 - res[i]
		if complement > resRemaining {
			complement = resRemaining
		}
		res[i] += complement
		resRemaining -= complement
	}

	return res
}
