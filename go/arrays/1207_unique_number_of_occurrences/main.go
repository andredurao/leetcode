package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	occurrencesMap := map[int]int{}

	for _, val := range arr {
		_, found := occurrencesMap[val]

		if !found {
			occurrencesMap[val] = 0
		}
		occurrencesMap[val]++
	}

	uniquenessMap := map[int]struct{}{}

	for _, count := range occurrencesMap {
		_, found := uniquenessMap[count]

		if found {
			return false
		}
		uniquenessMap[count] = struct{}{}
	}

	return true
}
