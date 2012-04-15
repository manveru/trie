// Ternary Search Tree
package trie

// An empty value of this type can be used to Put values.
type Trie struct {
	root *node
}

type node struct {
	small *node
	equal *node
	large *node
	value interface{}
	r     rune
	end   bool
}

// Put adds a key/value pair to the Trie.
//
// Complexity: O(m)
func (t *Trie) Put(key string, value interface{}) {
	if len(key) < 1 {
		return
	}

	t.root = t.putRecursive(t.root, []rune(key), 0, value)
}

// Get returns the value of the desired key, or nil if the key wasn't found.
//
// Complexity: O(m) worst case
func (t *Trie) Get(key string) interface{} {
	n := t.root

	for index, r := range key {
	Start:
		if n == nil {
			return nil
		}

		if r < n.r {
			n = n.small
			goto Start
		} else if r > n.r {
			n = n.large
			goto Start
		} else if index < (len(key) - 1) {
			n = n.equal
		} else if n.end {
			return n.value
		}
	}

	return nil
}

// Wildcard Returns a sorted slice with matches for the key.
// The wildcard characters that match any character are '*' and '.'.
// If no match is found, an empty slice is returned.
//
// Complexity: O(n) worst case
func (t *Trie) Wildcard(key string) []string {
	if len(key) < 1 {
		return nil
	}

	return t.wildcardRecursive(t.root, []rune(key), 0, "")
}

// LongestPrefix returns the longest key that has a prefix in common with the key.
// If no match is found, "" is returned.
//
// Complexity: O(m) worst case
func (t *Trie) LongestPrefix(key string) string {
	if len(key) < 1 {
		return ""
	}

	length := t.prefixRecursive(t.root, []rune(key), 0)
	return key[0:length]
}

func (t *Trie) prefixRecursive(n *node, key []rune, index int) int {
	if n == nil || index == len(key) {
		return 0
	}

	length := 0
	recLen := 0
	r := key[index]

	if r < n.r {
		recLen = t.prefixRecursive(n.small, key, index)
	} else if r > n.r {
		recLen = t.prefixRecursive(n.large, key, index)
	} else {
		if n.end {
			length = index + 1
		}
		recLen = t.prefixRecursive(n.equal, key, index+1)
	}
	if length > recLen {
		return length
	}
	return recLen
}

func (t *Trie) wildcardRecursive(n *node, key []rune, index int, prefix string) (matches []string) {
	if n == nil || index == len(key) {
		return matches
	}

	r := key[index]
	isWild := r == '*' || r == '.'

	if isWild || r < n.r {
		matches = append(matches, t.wildcardRecursive(n.small, key, index, prefix)...)
	}
	if isWild || r > n.r {
		matches = append(matches, t.wildcardRecursive(n.large, key, index, prefix)...)
	}
	if isWild || r == n.r {
		newPrefix := prefix + string(n.r)
		if n.end {
			matches = append(matches, newPrefix)
		}
		matches = append(matches, t.wildcardRecursive(n.equal, key, index+1, newPrefix)...)
	}

	return matches
}

func (t *Trie) getRecursive(n *node, key []rune, index int) *node {
	if n == nil {
		return nil
	}

	r := key[index]
	if r < n.r {
		return t.getRecursive(n.small, key, index)
	} else if r > n.r {
		return t.getRecursive(n.large, key, index)
	} else if index < (len(key) - 1) {
		return t.getRecursive(n.equal, key, index+1)
	}

	if n.end {
		return n
	}

	return nil
}

func (t *Trie) putRecursive(n *node, key []rune, index int, value interface{}) *node {
	r := key[index]
	if n == nil {
		n = &node{r: r}
	}
	if r < n.r {
		n.small = t.putRecursive(n.small, key, index, value)
	} else if r > n.r {
		n.large = t.putRecursive(n.large, key, index, value)
	} else if index < (len(key) - 1) {
		n.equal = t.putRecursive(n.equal, key, index+1, value)
	} else {
		n.end = true
		n.value = value
	}

	return n
}
