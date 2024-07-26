//

package main

import (
	"fmt"
	"math"
)

func main() {
	// n := 4
	// edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	// distanceThreshold := 4
	// res := findTheCity(n, edges, distanceThreshold)
	// fmt.Println(res)

	// n := 5
	// edges := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
	// distanceThreshold := 2
	// res := findTheCity(n, edges, distanceThreshold)
	// fmt.Println(res)

	// n := 6
	// edges := [][]int{{0, 1, 10}, {0, 2, 1}, {2, 3, 1}, {1, 3, 1}, {1, 4, 1}, {4, 5, 10}}
	// distanceThreshold := 20
	// res := findTheCity(n, edges, distanceThreshold)
	// fmt.Println(res)

	n := 39
	edges := [][]int{{32, 33, 6066}, {9, 24, 4482}, {12, 23, 1781}, {6, 25, 1897}, {7, 15, 8633}, {12, 16, 2890}, {1, 30, 349}, {30, 31, 9738}, {11, 33, 9791}, {12, 34, 2418}, {18, 21, 4112}, {25, 29, 7258}, {1, 3, 4596}, {1, 8, 2224}, {8, 17, 9142}, {13, 23, 6498}, {29, 38, 9590}, {6, 28, 6956}, {4, 31, 9774}, {2, 30, 3967}, {6, 19, 8528}, {11, 13, 3068}, {2, 36, 2987}, {29, 37, 5395}, {14, 21, 5175}, {2, 4, 3214}, {17, 29, 196}, {9, 20, 4655}, {19, 36, 9637}, {15, 25, 1418}, {6, 33, 5843}, {22, 27, 2500}, {13, 34, 2553}, {0, 16, 1409}, {20, 30, 795}, {5, 34, 8623}, {9, 33, 2352}, {21, 29, 525}, {11, 30, 1720}, {14, 17, 7672}, {2, 34, 8525}, {3, 29, 6520}, {26, 29, 847}, {14, 18, 1323}, {27, 33, 2360}, {14, 23, 4009}, {21, 37, 7194}, {14, 38, 7686}, {2, 25, 8244}, {3, 21, 7009}, {20, 27, 8794}, {4, 32, 1865}, {14, 20, 3548}, {2, 3, 6502}, {21, 28, 1577}, {9, 15, 1030}, {24, 32, 5566}, {3, 5, 4979}, {18, 26, 4109}, {25, 33, 6545}, {12, 36, 5506}, {5, 33, 564}, {13, 22, 691}, {8, 13, 1955}, {18, 19, 4031}, {15, 37, 841}, {7, 27, 318}, {1, 25, 1626}, {15, 18, 7242}, {11, 12, 1446}, {24, 26, 725}, {5, 24, 7100}, {7, 37, 9453}, {20, 26, 2597}, {2, 10, 6982}, {19, 25, 1081}, {1, 35, 7350}, {4, 37, 8618}, {4, 17, 3751}, {16, 38, 1582}, {8, 15, 2040}, {18, 36, 3113}, {2, 11, 4287}, {13, 28, 3813}, {0, 32, 4375}, {3, 33, 5513}, {19, 26, 244}, {11, 23, 2454}, {16, 28, 3209}, {3, 34, 7579}, {2, 24, 6368}, {10, 25, 6483}, {8, 22, 5691}, {7, 19, 4154}, {17, 23, 8757}, {7, 11, 1931}, {4, 19, 7856}, {22, 32, 8456}, {2, 12, 2615}, {29, 36, 4506}, {14, 37, 9937}, {11, 27, 4164}, {26, 38, 7275}, {6, 11, 9853}, {3, 31, 9498}, {6, 27, 835}, {6, 35, 9750}, {14, 28, 2564}, {8, 21, 2069}, {3, 38, 6068}, {0, 25, 2793}, {4, 23, 5182}, {15, 36, 6692}, {18, 25, 8000}, {12, 31, 8724}, {15, 27, 146}, {1, 7, 6611}, {1, 36, 5780}, {9, 23, 5532}, {20, 28, 3097}, {30, 38, 108}, {15, 17, 7243}, {6, 36, 2094}, {32, 34, 6015}, {11, 26, 5442}, {16, 17, 1454}, {18, 35, 5012}, {28, 38, 73}, {0, 38, 5039}, {17, 33, 8088}, {33, 35, 1675}, {10, 38, 2895}, {29, 31, 1275}, {13, 38, 7541}, {13, 17, 3776}, {13, 26, 3980}, {0, 22, 5068}, {5, 14, 420}, {11, 38, 3823}, {24, 37, 6245}, {7, 18, 745}, {11, 22, 894}, {14, 19, 7170}, {0, 15, 7181}, {10, 18, 5059}, {0, 20, 2448}, {8, 33, 9989}, {28, 30, 5110}, {6, 20, 8021}, {5, 15, 4099}, {3, 37, 1375}, {8, 29, 2438}, {5, 27, 3915}, {16, 37, 1430}, {10, 30, 5871}, {8, 9, 4053}, {23, 24, 2305}, {23, 30, 1723}, {11, 35, 43}, {23, 25, 377}, {11, 28, 949}, {2, 27, 2637}, {26, 36, 1856}, {9, 25, 994}, {7, 8, 9375}, {19, 24, 9937}, {6, 23, 1727}, {3, 10, 6053}, {22, 28, 9815}, {12, 24, 1033}, {17, 30, 1795}, {2, 23, 9458}, {0, 34, 4091}, {21, 34, 8096}, {1, 18, 1031}, {20, 34, 944}, {2, 5, 4024}, {0, 24, 285}, {11, 20, 8137}, {22, 24, 4782}, {0, 17, 8309}, {15, 28, 3969}, {15, 21, 2276}, {31, 34, 5448}, {10, 34, 6433}, {1, 31, 1736}, {10, 16, 8362}, {16, 22, 4084}, {2, 6, 7867}, {7, 32, 1865}, {2, 16, 3438}, {11, 16, 1160}, {8, 32, 3509}, {6, 9, 1658}, {5, 19, 2762}, {0, 5, 4162}, {19, 30, 2333}, {3, 16, 3306}, {25, 27, 3425}, {22, 23, 8181}, {9, 18, 3861}, {16, 34, 7057}, {14, 34, 9239}, {9, 16, 1192}, {16, 32, 8649}, {23, 28, 2251}, {10, 37, 9831}, {4, 36, 1830}, {0, 28, 4997}, {35, 36, 1370}, {21, 38, 1609}, {4, 18, 2630}, {5, 20, 8504}, {10, 22, 1379}, {26, 35, 9343}, {16, 18, 2038}, {10, 23, 491}, {24, 38, 6111}, {35, 38, 8084}, {8, 20, 7034}}
	distanceThreshold := 6586
	res := findTheCity(n, edges, distanceThreshold)
	fmt.Println(res)

}

func debug(m *[][]int) {
	for _, row := range *m {
		for _, v := range row {
			fmt.Printf("%d\t", v)
		}
		fmt.Printf("\n")
	}
}

// ---------------------------------------------------------------

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	distances := floydMarshall(n, &edges)

	counts := map[int][]int{}
	minCount := 999
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if distances[i][j] <= distanceThreshold {
				count++
			}
		}
		counts[count] = append(counts[count], i)
		minCount = min(minCount, count)
	}

	res := counts[minCount][0]
	for i := range counts[minCount] {
		res = max(res, counts[minCount][i])
	}

	return res
}

// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
func floydMarshall(n int, edges *[][]int) [][]int {
	distances := make([][]int, n)

	for i := range distances {
		distances[i] = make([]int, n)
		if i == 0 {
			for j := range distances[i] {
				distances[i][j] = math.MaxInt32
			}
		} else {
			copy(distances[i], distances[i-1])
		}
	}
	for i := range distances {
		distances[i][i] = 0
	}

	for _, edge := range *edges {
		distances[edge[0]][edge[1]] = edge[2]
		distances[edge[1]][edge[0]] = edge[2]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				distances[i][j] = min(distances[i][j], distances[i][k]+distances[k][j])
			}
		}
	}

	return distances
}
