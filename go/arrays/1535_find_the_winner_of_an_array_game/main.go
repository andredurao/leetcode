// https://leetcode.com/problems/find-the-winner-of-an-array-game/description/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// arr := []int{2, 1, 3, 5, 4, 6, 7}
	// k := 2

	arr := []int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	k := 1000000000

	result := getWinner(arr, k)

	fmt.Println(result)
}

func getWinner(arr []int, k int) int {
	i := 0
	count := 0
	winner := -1
	l := createList(&arr)
	for count < k {
		if i == l.Len() {
			return winner
		}

		curWinner := fight(l)
		if curWinner != winner {
			count = 1
			winner = curWinner
		} else {
			count++
		}
		i++
	}

	return winner
}

func createList(arr *[]int) *list.List {
	l := list.New()
	for _, val := range *arr {
		l.PushBack(val)
	}
	return l
}

func fight(l *list.List) int {
	v0 := l.Front().Value.(int)
	v1 := l.Front().Next().Value.(int)
	winner := -1
	looser := -1
	var e *list.Element

	if v0 >= v1 {
		winner = v0
		looser = v1
		e = l.Front().Next()
	} else {
		winner = v1
		looser = v0
		e = l.Front()
	}

	// send the value to the end of the list
	l.Remove(e)
	l.PushBack(looser)

	return winner
}

func debug(l *list.List) {
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Printf("\n")
}
