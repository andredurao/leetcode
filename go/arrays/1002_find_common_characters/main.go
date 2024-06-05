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
		for i, count := range cur {
			if prev[i] == 0 {
				cur[i] = 0
				continue
			}
			if count > prev[i] {
				cur[i] = prev[i]
			}
		}
	}

	res := []string{}
	for ch, count := range cur {
		for i := byte(0); i < count; i++ {
			res = append(res, string(ch+'a'))
		}
	}
	return res
}

func wordFreq(word *string) []byte {
	f := make([]byte, 26)

	for _, ch := range *word {
		f[ch-'a']++
	}

	return f
}
