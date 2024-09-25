// https://leetcode.com/problems/sum-of-prefix-scores-of-strings

package main

import "fmt"

func main() {
	words := []string{"abc", "ab", "bc", "b"}
	// words := []string{"abcd"}
	res := sumPrefixScores(words)
	fmt.Println(res)
}

// ==================================================

func sumPrefixScores(words []string) []int {
	trie := NewTrie()

	for _, word := range words {
		trie.Insert(word)
	}

	res := make([]int, len(words))
	for i, word := range words {
		res[i] = trie.Search(word)
	}

	return res
}

// using trie from challenge #208

type Node struct {
	Children [26]*Node
	Count    int
}

type Trie struct {
	Root *Node
}

func NewTrie() Trie {
	return Trie{&Node{}}
}

func (this *Trie) Insert(word string) {
	node := this.Root
	for _, ch := range word {
		if node.Children[ch-'a'] == nil {
			node.Children[ch-'a'] = &Node{}
		}
		node = node.Children[ch-'a']
		node.Count++
	}
}

func (this *Trie) Search(word string) int {
	node := this.Root
	res := 0
	for _, ch := range word {
		node = node.Children[ch-'a']
		res += node.Count
	}
	return res
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.Root
	for _, ch := range prefix {
		if node.Children[ch-'a'] == nil {
			return false
		}
		node = node.Children[ch-'a']
	}
	return true
}
