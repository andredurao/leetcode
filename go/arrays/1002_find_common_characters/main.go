// https://leetcode.com/problems/find-common-characters/

package main

import "fmt"

func main() {
	words := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(words))
}

func commonChars(words []string) []string {
	cur := wordFreq(&words[0])

	for i, word := range words {
		if i == 0 {
			continue
		}
		prev := cur
		cur = wordFreq(&word)
		// filter out chars that were not repeated or with different frequencies
		for ch, count := range cur {
			if _, found := prev[ch]; !found {
				delete(cur, ch)
				continue
			}
			if count > prev[ch] {
				cur[ch] = prev[ch]
			}
		}
	}

	res := []string{}
	for ch, count := range cur {
		for i := 0; i < count; i++ {
			res = append(res, string(ch))
		}
	}
	return res
}

func wordFreq(word *string) map[rune]int {
	f := map[rune]int{}

	for _, ch := range *word {
		if _, found := f[ch]; !found {
			f[ch] = 0
		}
		f[ch]++
	}

	return f
}
