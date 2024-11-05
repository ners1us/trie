package trie

import (
	"sync"
)

// Trie represents the root of the trie data structure.
type Trie struct {
	root *node        // Root node of the trie
	mu   sync.RWMutex // Mutex to ensure concurrency safety
}

// node represents a single node in the trie.
type node struct {
	children [26]*node // Use pointers to nodes
	isEnd    bool      // Flag to indicate the end of a word
}

// NewTrie initializes a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: &node{},
	}
}

// Insert adds a word to the trie.
func (t *Trie) Insert(word string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if index < 0 || index >= 26 {
			continue
		}

		if current.children[index] == nil {
			current.children[index] = &node{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

// Search checks if a word exists in the trie.
func (t *Trie) Search(word string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if index < 0 || index >= 26 {
			return false
		}

		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.isEnd
}

// StartsWith checks if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	current := t.root
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if index < 0 || index >= 26 {
			return false
		}

		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return true
}
