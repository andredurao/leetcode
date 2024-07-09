//

package main

func main() {
	customers := [][]int{{1, 2}, {2, 5}, {4, 3}}
	res := averageWaitingTime(customers)
	println(res)
}

func averageWaitingTime(customers [][]int) float64 {
	totalTime := 0
	time := 0
	for _, customer := range customers {
		if customer[0] > time {
			time = customer[0]
		}
		finish := time + customer[1]
		time = finish
		totalTime += finish - customer[0]
	}

	return float64(totalTime) / float64(len(customers))
}
