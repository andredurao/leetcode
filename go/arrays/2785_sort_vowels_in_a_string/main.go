package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "lEetcOde"
	fmt.Println(sortVowels(s))
}

func sortVowels(s string) string {
	vowels := make([]int, 0)
	var b strings.Builder
	var result strings.Builder
	for _, c := range s {
		tmpC := c
		if tmpC >= 'a' && tmpC <= 'z' {
			tmpC = tmpC - 32
		}
		if tmpC == 'A' || tmpC == 'E' || tmpC == 'I' || tmpC == 'O' || tmpC == 'U' {
			vowels = append(vowels, int(c))
			b.WriteString("_")
		} else {
			b.WriteString(string(c))
		}
	}
	sort.IntSlice.Sort(vowels)

	i := 0
	for _, c := range b.String() {
		char := string(rune(c))
		if char == "_" {
			result.WriteString(string(rune(vowels[i])))
			i++
		} else {
			result.WriteString(char)
		}
	}

	return result.String()
}
