// https://leetcode.com/problems/walking-robot-simulation/

package main

import "fmt"

func main() {
	// commands := []int{4, -1, 3}
	// obstacles := [][]int{}

	// commands := []int{4, -1, 4, -2, 4}
	// obstacles := [][]int{{2, 4}}

	// commands := []int{-2, 8, 3, 7, -1}
	// obstacles := [][]int{{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}}

	commands := []int{-2, -1, -2, 3, 7} // (N) -> W, N, W, 3
	obstacles := [][]int{{1, -3}, {2, -3}, {4, 0}, {-2, 5}, {-5, 2}, {0, 0}, {4, -4}, {-2, -5}, {-1, -2}, {0, 2}}
	res := robotSim(commands, obstacles)
	fmt.Println(res)
}

// --------------------

func robotSim(commands []int, obstacles [][]int) int {
	obsMap := map[int]map[int]struct{}{}
	for _, obs := range obstacles {
		_, found := obsMap[obs[0]]
		if !found {
			obsMap[obs[0]] = map[int]struct{}{}
		}
		obsMap[obs[0]][obs[1]] = struct{}{}
	}

	pos := []int{0, 0}
	dir := 0
	res := 0

	for _, cmd := range commands {
		// fmt.Println("cmd", cmd, "dir", dir, "pos", pos)
		if cmd > 0 {
			pos = walk(pos, cmd, dir, obsMap)
		} else {
			changeDir(&dir, cmd)
		}
		res = max(res, pos[0]*pos[0]+pos[1]*pos[1])
	}
	// fmt.Println("end pos", pos)

	return res
}

func walk(pos []int, cmd, dir int, obsMap map[int]map[int]struct{}) []int {
	// dirs -> N E S W
	if dir == 0 { // N
		goal := pos[1] + cmd
		init := pos[1]
		for ; pos[1] < goal; pos[1]++ {
			if isObstacle(pos, obsMap) && pos[1] != init {
				pos[1]--
				// fmt.Println("obs", pos)
				return pos
			}
		}
	} else if dir == 1 { // E
		goal := pos[0] + cmd
		init := pos[0]
		for ; pos[0] < goal; pos[0]++ {
			if isObstacle(pos, obsMap) && pos[0] != init {
				pos[0]--
				// fmt.Println("obs", pos)
				return pos
			}
		}
	} else if dir == 2 { // S
		goal := pos[1] - cmd
		init := pos[1]
		for ; pos[1] > goal; pos[1]-- {
			if isObstacle(pos, obsMap) && pos[1] != init {
				pos[1]++
				// fmt.Println("obs", pos)
				return pos
			}
		}
	} else if dir == 3 { // W
		goal := pos[0] - cmd
		init := pos[0]
		for ; pos[0] > goal; pos[0]-- {
			if isObstacle(pos, obsMap) && pos[0] != init {
				pos[0]++
				// fmt.Println("obs", pos)
				return pos
			}
		}
	}

	return pos
}

func isObstacle(pos []int, obsMap map[int]map[int]struct{}) bool {
	if _, found := obsMap[pos[0]]; found {
		if _, found := obsMap[pos[0]][pos[1]]; found {
			return true
		}
	}
	return false
}

func changeDir(dir *int, command int) {
	// dirs -> N E S W
	if command == -1 { // turn right 90 degrees
		*dir = (*dir + 1) % 4
	} else if command == -2 {
		*dir = (*dir + 3) % 4
	}
}
