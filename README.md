# trie

Module with concurrent safe trie implementation.

## Installation

```bash
go get github.com/ners1us/trie
```

## What is a Trie?

A **trie**, also known as a **prefix tree**, is a specialized tree data structure used to store a dynamic set of strings. Unlike binary trees, tries are particularly efficient for retrieval operations, especially when dealing with prefixes.

### Key Characteristics

- **Nodes and Edges**: Each node represents a single character of a string, and edges connect these characters to form words.
- **Root Node**: The trie starts with an empty root node that doesn't hold any character.
- **End of Word Marker**: Nodes can be marked to indicate the completion of a valid word.
- **Shared Prefixes**: Common prefixes are shared among multiple words, optimizing space.

This implementation uses `RWMutex` and prevents race conditions by synchronizing concurrent read and write access to the trie.

## Core Methods

### Insert(word string)

The `Insert` method adds a new word to the trie. Here's how it works:

1. Starts from the root node
2. For each character in the word:
   - Calculates array index (0-25) by subtracting 'a' from the character
   - If the index is invalid (not a-z), skips the character
   - If no node exists at the calculated index, creates a new one
   - Moves to the created/existing node
3. Marks the final node as end of word (isEnd = true)

### Search(word string) bool

The `Search` method checks if a complete word exists in the trie:

1. Starts from the root node
2. For each character in the word:
   - Calculates array index (0-25)
   - Returns false if:
      - Index is invalid (not a-z)
      - No node exists at the calculated index
   - Moves to the next node if found
3. Returns true only if:
   - All characters were found
   - Final node is marked as end of word (isEnd = true)

### StartsWith(prefix string) bool

The `StartsWith` method checks if any word in the trie begins with the given prefix:

1. Starts from the root node
2. For each character in the prefix:
   - Calculates array index (0-25)
   - Returns false if:
      - Index is invalid (not a-z)
      - No node exists at the calculated index
   - Moves to the next node if found
3. Returns true if all prefix characters were found

Unlike `Search`, this method doesn't check the `isEnd` flag since it only verifies the prefix existence.

## Example

Consider inserting the words: cat, car and dog into a trie:

```
root
├── c
│   └── a
│       ├── t (End of Word)
│       └── r (End of Word)
└── d
    └── o
        └── g (End of Word)
```

### Insertion Steps

1. **Insert "cat"**
   - Create nodes: c → a → t
   - Mark 't' as the end of a word

2. **Insert "car"**
   - Traverse existing nodes: c → a
   - Add node 'r' after 'a'
   - Mark 'r' as the end of a word

3. **Insert "dog"**
   - Add a new branch: d → o → g
   - Mark 'g' as the end of a word