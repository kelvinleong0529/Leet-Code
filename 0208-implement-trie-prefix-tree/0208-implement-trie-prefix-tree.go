// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

// Constructor initializes the Trie object
func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert inserts a word into the Trie
func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search returns true if the word is in the Trie
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

// StartsWith returns true if there is a word in the Trie that starts with the given prefix
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
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