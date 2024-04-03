//

package main

import (
	"fmt"
	"os"
)

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := os.Args[1]
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%s\t", string(board[i][j]))
		}
		fmt.Printf("\n")
	}
	fmt.Println("looking for", word)
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			ch := board[i][j]
			if ch == word[0] {
				if search(&board, &word, i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

func search(board *[][]byte, word *string, i int, j int, wordIndex int) bool {
	// word found
	if wordIndex == len(*word) {
		return true
	}

	// check if i and j are out of bounds
	if i < 0 || i >= len(*board) || j < 0 || j >= len((*board)[0]) {
		return false
	}

	// check if current char is correct
	expectedChar := (*word)[wordIndex]
	if (*board)[i][j] != expectedChar {
		return false
	}

	// mark that char as visited
	(*board)[i][j] = '-'

	for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		if search(board, word, i+dir[0], j+dir[1], wordIndex+1) {
			return true
		}
	}

	// restore visited value
	(*board)[i][j] = expectedChar
	return false
}
