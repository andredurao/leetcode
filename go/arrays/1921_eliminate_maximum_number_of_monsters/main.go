package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// dist := []int{1, 3, 4}
	// speed := []int{1, 1, 1}
	dist := []int{4, 2, 3}
	speed := []int{2, 1, 1}
	result := eliminateMaximum(dist, speed)
	fmt.Println(result)
}

func eliminateMaximum(dist []int, speed []int) int {
	etas := make([][]int, len(dist))
	for i := range dist {
		etas[i] = []int{i, eta(&dist, &speed, i)}
	}
	sort.Slice(etas, func(i, j int) bool { return etas[i][1] < etas[j][1] })
	fmt.Println(etas)

	total := 0
	time := 0
	i := 0
	for i < len(etas) {
		cur := etas[i]
		fmt.Println(i, cur, time)
		if cur[1] <= time {
			break
		}
		i++
		total++
		time++
	}

	return total
}

func eta(dist *[]int, speed *[]int, index int) int {
	d := float64((*dist)[index])
	s := float64((*speed)[index])
	return int(math.Ceil(d / s))
}



dist = [4,2,3] ; speed = [2,1,1]
0   4 2 3
1   2 1 2
2   0 0 1


dist = [3,2,4], speed = [5,3,2]
0 3 2 4
1 x
