// https://leetcode.com/problems/reverse-prefix-of-word/description/

package main

func main() {
	println(reversePrefix("abcdefd", 'd'))
}

func reversePrefix(word string, ch byte) string {
	prefixIndex := -1
	for prefixIndex = range word {
		if word[prefixIndex] == ch {
			break
		}
	}
	if prefixIndex == len(word)-1 && word[len(word)-1] != ch {
		return word
	}
	prefix := make([]byte, prefixIndex+1)

	for i := prefixIndex; i >= 0; i-- {
		prefix[len(prefix)-i-1] = word[i]
	}
	return string(prefix) + word[(prefixIndex+1):]
}
