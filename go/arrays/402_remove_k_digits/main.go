package main

import (
	"bytes"
	"container/list"
	"fmt"
)

func main() {
	num := "112"
	// res := getMinNumber(num)
	// fmt.Println(res)
	res := removeKdigits(num, 1)
	fmt.Println(res)
}

// from Editoral
func removeKdigits(num string, k int) string {
	ll := list.New()

	for _, ch := range num {
		for k > 0 && ll.Len() > 0 && ll.Back().Value.(rune) > ch {
			ll.Remove(ll.Back())
			k--
		}
		ll.PushBack(ch)
	}

	for i := 0; i < k; i++ {
		ll.Remove(ll.Back())
	}

	var buffer bytes.Buffer
	leadingZero := true

	for e := ll.Front(); e != nil; e = e.Next() {
		ch := e.Value.(rune)
		if leadingZero && ch == '0' {
			continue
		}
		leadingZero = false
		buffer.WriteRune(ch)
	}

	if buffer.Len() == 0 {
		buffer.WriteRune('0')
	}
	return buffer.String()
}
