package trie

const AlphabetSize = 26

type TrieNode struct {
	children [AlphabetSize]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(w string) {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &TrieNode{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func InitTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}
