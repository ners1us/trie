package trie

import (
	"sync"
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	trie := NewTrie()

	// Test inserting and searching for words
	trie.Insert("hello")
	trie.Insert("world")

	if !trie.Search("hello") {
		t.Error("Expected 'hello' to be found in the trie")
	}

	if !trie.Search("world") {
		t.Error("Expected 'world' to be found in the trie")
	}

	if trie.Search("hell") {
		t.Error("Expected 'hell' to not be found in the trie")
	}

	if trie.Search("worlds") {
		t.Error("Expected 'worlds' to not be found in the trie")
	}
}

func TestStartsWith(t *testing.T) {
	trie := NewTrie()

	trie.Insert("apple")
	trie.Insert("app")

	// Test prefix search
	if !trie.StartsWith("app") {
		t.Error("Expected prefix 'app' to be found in the trie")
	}

	if trie.StartsWith("applz") {
		t.Error("Expected prefix 'applz' to not be found in the trie")
	}

	if trie.StartsWith("b") {
		t.Error("Expected prefix 'b' to not be found in the trie")
	}
}

func TestRemove(t *testing.T) {
	trie := NewTrie()

	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("apply")

	// Ensure all words are present initially
	if !trie.Search("apple") {
		t.Error("Expected 'apple' to be found in the trie")
	}
	if !trie.Search("app") {
		t.Error("Expected 'app' to be found in the trie")
	}
	if !trie.Search("apply") {
		t.Error("Expected 'apply' to be found in the trie")
	}

	// Remove "apple" and ensure that only it is removed
	trie.Remove("apple")
	if trie.Search("apple") {
		t.Error("Expected 'apple' to be removed from the trie")
	}
	if !trie.Search("app") {
		t.Error("Expected 'app' to still be found in the trie")
	}
	if !trie.Search("apply") {
		t.Error("Expected 'apply' to still be found in the trie")
	}

	// Remove "app" and ensure that it doesn't affect "apply"
	trie.Remove("app")
	if trie.Search("app") {
		t.Error("Expected 'app' to be removed from the trie")
	}
	if !trie.Search("apply") {
		t.Error("Expected 'apply' to still be found in the trie")
	}

	// Remove "apply" and ensure that trie is empty
	trie.Remove("apply")
	if trie.Search("apply") {
		t.Error("Expected 'apply' to be removed from the trie")
	}

	trie.Remove("nonexistent")
	if trie.Search("app") || trie.Search("apple") || trie.Search("apply") {
		t.Error("Trie state changed incorrectly after trying to remove a non-existent word")
	}
}

func TestInsertAndSearchWithNonAlphabeticChars(t *testing.T) {
	trie := NewTrie()

	trie.Insert("hello123")
	trie.Insert("world!")
	trie.Insert("hello")

	// Check that only valid words are found
	if !trie.Search("hello") {
		t.Error("Expected 'hello' to be found in the trie")
	}

	if trie.Search("hello123") {
		t.Error("Expected 'hello123' to not be found in the trie")
	}

	if trie.Search("world!") {
		t.Error("Expected 'world!' to not be found in the trie")
	}
}

func TestConcurrentInsertAndSearch(t *testing.T) {
	trie := NewTrie()
	var wg sync.WaitGroup

	// Concurrently insert words into the Trie
	wordsToInsert := []string{"concurrent", "safe", "trie", "test", "data"}
	for _, word := range wordsToInsert {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			trie.Insert(word)
		}(word)
	}

	wg.Wait()

	// Concurrently search for words in the Trie
	for _, word := range wordsToInsert {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			if !trie.Search(word) {
				t.Errorf("Expected '%s' to be found in the trie", word)
			}
		}(word)
	}

	// Concurrently check prefixes
	prefixesToCheck := []string{"con", "sa", "tri", "tes", "da"}
	for _, prefix := range prefixesToCheck {
		wg.Add(1)
		go func(prefix string) {
			defer wg.Done()
			if !trie.StartsWith(prefix) {
				t.Errorf("Expected prefix '%s' to be found in the trie", prefix)
			}
		}(prefix)
	}

	wg.Wait()
}

func TestConcurrentInsertAndFailingSearch(t *testing.T) {
	trie := NewTrie()
	var wg sync.WaitGroup

	// Insert words concurrently
	wordsToInsert := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	for _, word := range wordsToInsert {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			trie.Insert(word)
		}(word)
	}

	wg.Wait()

	// Search for a word that wasn't inserted, expecting failure
	if trie.Search("omega") {
		t.Error("Expected 'omega' to not be found in the trie, but it was found")
	}

	// Concurrently search for inserted words to check concurrency handling
	for _, word := range wordsToInsert {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			if !trie.Search(word) {
				t.Errorf("Expected '%s' to be found in the trie, but it wasn't", word)
			}
		}(word)
	}

	// Add a search for a non-existing word concurrently to cause a fail due to concurrency issues
	wg.Add(1)
	go func() {
		defer wg.Done()
		if trie.Search("nonexistent") {
			t.Error("Expected 'nonexistent' to not be found in the trie, but it was found")
		}
	}()

	wg.Wait()
}

func BenchmarkTrieInsert(b *testing.B) {
	tr := NewTrie()
	words := []string{"apple", "banana", "grape", "orange", "watermelon"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		word := words[i%len(words)]
		tr.Insert(word)
	}
}

func BenchmarkTrieSearch(b *testing.B) {
	tr := NewTrie()
	words := []string{"apple", "banana", "grape", "orange", "watermelon"}
	for _, word := range words {
		tr.Insert(word)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		word := words[i%len(words)]
		tr.Search(word)
	}
}

func BenchmarkTrieStartsWith(b *testing.B) {
	tr := NewTrie()
	words := []string{"apple", "banana", "grape", "orange", "watermelon"}
	for _, word := range words {
		tr.Insert(word)
	}

	prefixes := []string{"a", "b", "gr", "ora", "w"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefix := prefixes[i%len(prefixes)]
		tr.StartsWith(prefix)
	}
}

func BenchmarkTrieRemove(b *testing.B) {
	tr := NewTrie()
	words := []string{"apple", "banana", "grape", "orange", "watermelon", "application", "apply", "app", "grapefruit"}

	for _, word := range words {
		tr.Insert(word)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		word := words[i%len(words)]
		tr.Remove(word)
		tr.Insert(word)
	}
}
