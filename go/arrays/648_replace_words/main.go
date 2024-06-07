// https://leetcode.com/problems/replace-words/

package main

import (
	"strings"
)

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	println(replaceWords(dictionary, sentence))
}

func replaceWords(dictionary []string, sentence string) string {
	trie := Trie{&Node{Terminal: false}}
	for _, root := range dictionary {
		trie.Insert(root)
	}

	res := []string{}
	for _, word := range strings.Split(sentence, " ") {
		prefix, found := trie.SearchPrefix(word)
		if found {
			res = append(res, prefix)
		} else {
			res = append(res, word)
		}
	}

	return strings.Join(res, " ")
}

// from problem 208
type Node struct {
	Children [26]*Node
	Terminal bool
}

type Trie struct {
	Root *Node
}

func (this *Trie) Insert(word string) {
	node := this.Root
	for _, ch := range word {
		// early return for substrings of the current prefix that have been already inserted
		// if node.Terminal {
		// 	return
		// }

		if node.Children[ch-'a'] == nil {
			node.Children[ch-'a'] = &Node{Terminal: false}
		}
		node = node.Children[ch-'a']
	}
	node.Terminal = true
}

func (this *Trie) SearchPrefix(word string) (string, bool) {
	node := this.Root
	prefix := ""
	for _, ch := range word {
		if node.Terminal {
			return prefix, true
		}

		if node.Children[ch-'a'] == nil {
			return "", false
		}
		prefix += string(ch)
		node = node.Children[ch-'a']
	}
	return prefix, true
}
