package main

import "fmt"

func main() {
	// s := "ab#c"
	// t := "ad#c"
	s := "a##c"
	t := "#a#c"
	result := backspaceCompare(s, t)
	fmt.Println(result)
}

func backspaceCompare(s string, t string) bool {
	cleanString(&s)
	cleanString(&t)
	return s == t
}

func cleanString(s *string) {
	i := 0
	for i < len(*s) {
		// fmt.Println("s", *s, "ch", (*s)[i], "i", i)
		if (*s)[i] == '#' {
			l_end := i - 1
			if l_end < 0 {
				l_end = 0
			}
			left := (*s)[0:l_end]

			right := ""
			r_start := i + 1
			if r_start < len(*s) {
				right = (*s)[r_start:]
			}

			// fmt.Println("ch", left, right)
			*s = left + right
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
}
