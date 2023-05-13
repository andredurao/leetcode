package main

import (
	"container/list"
	"fmt"
)

func main() {
	colors := "AAABABB"
	result := winnerOfGame(colors)
	fmt.Printf("%v\n", result)
}

func winnerOfGame(colors string) bool {
	if len(colors) < 3 {
		return false
	}

	linked_list := init_list(colors)
	fmt.Printf("colors %s\n", colors)
	for linked_list.Len() >= 3 {
		fmt.Println("ALICE")
		if play(linked_list, 'A') {
			fmt.Println("ALICE removed a letter")
			debug(linked_list)
		} else {
			fmt.Println("BOB WINS")
			return false
		}

		fmt.Println("BOB")
		if play(linked_list, 'B') {
			fmt.Println("BOB removed a letter")
			debug(linked_list)
		} else {
			fmt.Println("ALICE WINS")
			return true
		}
	}

	return false
}

func init_list(colors string) *list.List {
	l := list.New()
	for _, char := range colors {
		l.PushBack(char)
	}
	return l
}

// given a color, return the index if the play was possible, otherwise -1
func play(list *list.List, color rune) bool {
	if list.Len() < 3 {
		return false
	}
	for e := list.Front().Next(); e.Next() != nil; e = e.Next() {
		if color == e.Value && color == e.Prev().Value && color == e.Next().Value {
			list.Remove(e)
			return true
		}
	}
	return false
}

func debug(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		char, _ := e.Value.(rune)
		fmt.Printf("%s", string(char))
	}
	fmt.Printf("\n")
}
