// https://leetcode.com/problems/get-equal-substrings-within-budget/

package main

func main() {
	s := "ujteygggjwxnfl"
	t := "nstsenrzttikoy"
	maxCost := 43

	println(equalSubstring(s, t, maxCost))
}

func equalSubstring(s string, t string, maxCost int) int {
	diffs := map[int]int{}

	for i := range s {
		if s[i] != t[i] {
			diff := int(int(s[i]) - int(t[i]))
			if diff < 0 {
				diff = -diff
			}
			diffs[i] = diff
		}
	}

	l := 0
	r := 0
	total := 0
	cost := 0
	for r < len(s) {
		if diff, found := diffs[r]; found {
			cost += diff
		}
		r++

		for cost > maxCost {
			if diff, found := diffs[l]; found {
				cost -= diff
			}
			l++
		}
		total = max(total, (r - l))
	}

	return total
}
