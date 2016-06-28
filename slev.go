package levenshtein

import (
	"math"
)

type Row []int32

func SLev(w1 string, w2 string) int32 {
	cols := len(w1) + 1
	rows := len(w2) + 1

	var currentRow Row
	currentRow = Row(make([]int32, cols))
	currentRow[0] = 0
	for i := 1; i < cols; i++ {
		currentRow[i] = currentRow[i-1] + 1
	}

	for r := 1; r < rows; r++ {
		previousRow := currentRow
		currentRow = Row(make([]int32, cols))
		currentRow[0] = previousRow[0] + 1
		var rc int32
		for c := 1; c < cols; c++ {
			ic := currentRow[c-1] + 1
			dc := previousRow[c] + 1
			if w1[c-1] != w2[r-1] {
				rc = previousRow[c-1] + 1
			} else {
				rc = previousRow[c-1]
			}
			currentRow[c] = int32(math.Min(float64(ic), math.Min(float64(dc), float64(rc))))
		}
	}
	return currentRow[cols-1]
}

// Simple trie

type Trie struct {
	root *trieNode
}

type trieNode struct {
	word     string
	children map[rune]*trieNode
}

func NewTrie() *Trie {
	t := new(Trie)
	t.root = newTrieNode()
	return t
}
func (t Trie) Insert(word string) {
	t.root.insert(t.root, word)
}

func (t Trie) Search(word string, cost int32) {
	results := make([]string, 0)
	currentRow := make([]int32, 0)
	for i := 0; i <= len(word); i++ {
		currentRow = append(currentRow, i)
	}

	for k, v := range t.root.children {
		searchRecursive(v, k, word, currentRow, results, cost)
	}
}

func searchRecursive(node *trieNode, letter rune, word string, previousRow []int32, results []string, cost int32) {
	columns := len(word) + 1
	currentRow := make([]int32, 0)
	currentRow = append(currentRow, previousRow[0]+1)

	for c := 1, c < columns, c++ {
		ic := currentRow[c - 1] + 1
		dc := previousRow[c] + 1
		var rc int32
		if word[c - 1] != letter {
			rc = previousRow[c - 1] + 1
		} else {
			rc = previousRow[ c - 1]
		}
		currentRow = append(currentRow, int32(math.Min(float64(ic), math.Min(float64(dc), float64(rc)))))
	}
	if currentRow[len(currentRow) - 1] <= cost && node.word != nil {
		results = append(results, node.word)
	}
}

func newTrieNode() *trieNode {
	n := new(trieNode)
	n.children = make(map[rune]*trieNode)
	return n
}

func (t trieNode) insert(self *trieNode, word string) {
	node := self
	for _, l := range word {
		if _, ok := t.children[l]; !ok {
			node.children[l] = newTrieNode()
		}
	}
	node.word = word
}
