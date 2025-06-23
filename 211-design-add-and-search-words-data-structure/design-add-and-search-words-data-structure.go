type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, r := range word {
		index := r - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {

	var dfs func(*TrieNode, int) bool
	dfs = func(node *TrieNode, index int) bool {

		if node == nil {
			return false
		}

		if index == len(word) {
			return node.isEnd
		}

        char := word[index]

		if char == '.' {
			for i := 0; i < 26; i++ {
				if dfs(node.children[i], index+1) {
					return true
				}
			}
			return false
		} else {
			child := node.children[char-'a']
			return dfs(child, index+1)
		}
	}

	return dfs(this.root, 0)
}