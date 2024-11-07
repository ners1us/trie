# trie

Module with concurrent safe trie implementation

[//]: # (## Installation)

[//]: # ()
[//]: # (```bash)

[//]: # (go get github.com/ners1us/trie)

[//]: # (```)

## What is a Trie?

A **trie**, also known as a **prefix tree**, is a specialized tree data structure used to store a dynamic set of strings. Unlike binary trees, tries are particularly efficient for retrieval operations, especially when dealing with prefixes.

### Key Characteristics

- **Nodes and Edges**: Each node represents a single character of a string, and edges connect these characters to form words.
- **Root Node**: The trie starts with an empty root node that doesn't hold any character.
- **End of Word Marker**: Nodes can be marked to indicate the completion of a valid word.
- **Shared Prefixes**: Common prefixes are shared among multiple words, optimizing space.

This implementation uses `RWMutex` and prevents race conditions by synchronizing concurrent read and write access to the trie.

## Example

Consider inserting the words: cat, car and dog into a trie.

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
    - Mark 't' as the end of a word.

2. **Insert "car"**
    - Traverse existing nodes: c → a
    - Add node 'r' after 'a'.
    - Mark 'r' as the end of a word.

3. **Insert "dog"**
    - Add a new branch: d → o → g
    - Mark 'g' as the end of a word.
