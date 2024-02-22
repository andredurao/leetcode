//

package main

func main() {
	n := 2
	trust := [][]int{{1, 2}}
	println(findJudge(n, trust))
}

func findJudge(n int, trust [][]int) int {
	// a map that contain for each n an array (a) where:
	// a[0] = qty of people on which n trusts
	// a[1] = qty of people that trusts in n
	trustsMap := map[int][]int{}

	// initialize the map
	for i := 1; i <= n; i++ {
		trustsMap[i] = []int{0, 0}
	}

	for _, pair := range trust {
		trustsMap[pair[0]][0]++
		trustsMap[pair[1]][1]++
	}

	for key, counts := range trustsMap {
		if counts[0] == 0 && counts[1] == n-1 {
			return key
		}
	}

	return -1
}
