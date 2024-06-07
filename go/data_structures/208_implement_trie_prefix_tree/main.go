package main

func main() {

	trie := Constructor()
	trie.Insert("apple")
	trie.Search("apple")
	trie.Search("app")
	trie.StartsWith("app")
	// trie.Insert("app")
	// trie.Search("app")
}

// ===========================

type Node struct {
	Children []*Node
	Terminal bool
	Val      string
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{&Node{make([]*Node, 26), false, ""}}
}

func (this *Trie) Insert(word string) {
	node := this.Root
	for _, ch := range word {
		if node.Children[ch-'a'] == nil {
			node.Children[ch-'a'] = &Node{make([]*Node, 26), false, ""}
		}
		node = node.Children[ch-'a']
	}
	node.Val = word
	node.Terminal = true
}

func (this *Trie) Search(word string) bool {
	node := this.Root
	for _, ch := range word {
		if node.Children[ch-'a'] == nil {
			return false
		}
		node = node.Children[ch-'a']
	}
	return node.Terminal
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

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
