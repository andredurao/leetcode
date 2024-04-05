// https://leetcode.com/problems/make-the-string-great/description/

package main

import (
	"fmt"
	"os"
)

func main() {
	// "leEeetcode" "abBAcC"
	fmt.Println(makeGood(os.Args[1]))
}

func makeGood(s string) string {
	queue := []rune{}
	for _, ch := range s {
		// for a single toUpper with no conditions: ((ch-97+128)%32 + 65) <- In go the `+128` is needed because go returns negative remainders for `%`
		// but in this case I'm interested in only aA or Aa conditions
		if len(queue) > 0 && ((queue[len(queue)-1]-ch == 32) || (ch-queue[len(queue)-1] == 32)) {
			queue = queue[:len(queue)-1]
		} else {
			queue = append(queue, ch)
		}
	}
	return string(queue)
}
