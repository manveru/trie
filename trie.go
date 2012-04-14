package trie

// naming chosen to align, not to be meaningful :P
type node struct {
	small *node
	equal *node
	large *node
	value interface{}
	r     rune
	end   bool
}

type Trie struct {
	root *node
}

func (t *Trie) Put(key string, value interface{}) {
	if len(key) < 1 {
		return
	}

	t.root = t.putRecursive(t.root, []rune(key), 0, value)
}

func (t *Trie) Get(key string) interface{} {
	if len(key) < 1 {
		return nil
	}

	node := t.getRecursive(t.root, []rune(key), 0)
	if node == nil {
		return nil
	}
	return node.value
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
