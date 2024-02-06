// https://leetcode.com/problems/group-anagrams/submissions/1167843101/?envType=daily-question&envId=2024-02-06

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for _, str := range strs {
		fmt.Println(str, sortedString(str))
	}
	res := groupAnagrams(strs)
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	strMap := map[string][]string{}

	for _, str := range strs {
		sortedStr := sortedString(str)
		_, found := strMap[sortedStr]
		if !found {
			strMap[sortedStr] = []string{}
		}
		strMap[sortedStr] = append(strMap[sortedStr], str)
	}

	for _, values := range strMap {
		res = append(res, values)
	}

	return res
}

func sortedString(str string) string {
	chars := []rune(str)
	sort.Slice(chars, func(i int, j int) bool { return chars[i] < chars[j] })
	res := ""
	for _, ch := range chars {
		res += string(ch)
	}
	return res
}
