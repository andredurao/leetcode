//

package main

func main() {
	s := "hello"
	println(scoreOfString(s))
}

func scoreOfString(s string) int {
	total := 0
	for i := 0; i < len(s)-1; i++ {
		diff := int(s[i]) - int(s[i+1])
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}
	return total
}
