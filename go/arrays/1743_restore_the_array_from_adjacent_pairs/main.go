// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/description/

package main

import "fmt"

func main() {
	// adjacentPairs := [][]int{{2, 1}, {3, 4}, {3, 2}}
	// adjacentPairs := [][]int{{4, -2}, {1, 4}, {-3, 1}}
	// adjacentPairs := [][]int{{100000, -100000}}
	// adjacentPairs := [][]int{{-3, -9}, {-5, 3}, {2, -9}, {6, -3}, {6, 1}, {5, 3}, {8, 5}, {-5, 1}, {7, 2}}
	// adjacentPairs := [][]int{{2, 1}, {3, 4}, {3, 2}}
	adjacentPairs := [][]int{
		{91207, 59631},
		{581, 91207},
		{51465, 20358},
		{-66119, 94118},
		{-7392, 72809},
		{51465, -7392},
		{-98504, -29411},
		{-98504, 35968},
		{35968, 59631},
		{94118, 20358},
		{72809, 581},
		{34348, 43653},
		{43653, -29411},
	}

	fmt.Println("input", adjacentPairs)
	fmt.Println("end", restoreArray(adjacentPairs))
}

func restoreArray(adjacentPairs [][]int) []int {
	if len(adjacentPairs) == 1 {
		return adjacentPairs[0]
	}

	pairsMap := make(map[int][][]int, 0)
	for _, pair := range adjacentPairs {
		addPair(&pair, 0, pairsMap)
		addPair(&pair, 1, pairsMap)
	}

	// for val, pairs := range pairsMap {
	// 	fmt.Printf("%d -> %v\n", val, pairs)
	// }

	// lookup for initial values
	prev := &adjacentPairs[0]
	result := make([]int, 2)
	for val, pair := range pairsMap {
		if len(pair) == 1 {
			prev = &pair[0]
			result[0] = val
			if pair[0][0] == val {
				result[1] = pair[0][1]
			} else {
				result[1] = pair[0][0]
			}
			break
		}
	}
	// fmt.Println("initial", result)

	for {
		lastValue := result[len(result)-1]
		if len(pairsMap[lastValue]) == 1 {
			// last pair
			return result
		} else {
			current := &pairsMap[lastValue][0]
			if cmp(current, prev) {
				current = &pairsMap[lastValue][1]
			}

			index := 0
			if (*current)[index] == lastValue {
				index = 1
			}
			result = append(result, (*current)[index])
			prev = current
		}
	}
}

func cmp(pair1 *[]int, pair2 *[]int) bool {
	return (*pair1)[0] == (*pair2)[0] && (*pair1)[1] == (*pair2)[1]
}

func addPair(pair *[]int, index int, pairsMap map[int][][]int) {
	_, found := pairsMap[(*pair)[index]]
	if !found {
		pairsMap[(*pair)[index]] = [][]int{*pair}
	} else {
		pairsMap[(*pair)[index]] = append(pairsMap[(*pair)[index]], *pair)
	}
}
