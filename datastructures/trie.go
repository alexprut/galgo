package datastructures

type Trie struct {
	root trienode
}

func NewTrie() Trie {
	return Trie{newTrieNode()}
}

func (trie *Trie) Insert(elem string) {
	if len(elem) > 0 {
		tmp := trie.root
		for i := 0; i < len(elem); i++ {
			tmp.size++
			if _, ok := tmp.children[string(elem[i])]; !ok {
				tmp.children[string(elem[i])] = newTrieNode()
			}
			tmp = tmp.children[string(elem[i])]
		}
		tmp.size++
		tmp.isEndOfWord = true
	}
}

func (trie *Trie) SearchPrefix(prefix string) bool {
	return trie.contains(prefix) != nil
}

func (trie *Trie) Search(elem string) bool {
	node := trie.contains(elem)
	return node != nil && node.isEndOfWord
}

func (trie *Trie) SizePrefix(prefix string) int {
	node := trie.contains(prefix)
	if node != nil {
		return node.size
	} else {
		return 0
	}
}

func (trie *Trie) Size() int {
	return trie.root.size
}

func (trie *Trie) Remove(elem string) bool {
	if len(elem) > 0 && trie.Search(elem) {
		tmp := trie.root
		for i := 0; i < len(elem); i++ {
			tmp.size--
			if e := tmp.children[string(elem[i])]; e.size > 1 {
				tmp = tmp.children[string(elem[i])]
			} else {
				delete(tmp.children, string(elem[i]))
				break
			}
		}
		tmp.size--
		return true
	}

	return false
}

func (trie *Trie) RemovePrefix(prefix string) bool {
	if len(prefix) > 0 && trie.SearchPrefix(prefix) {
		trie.root.size -= trie.root.children[string(prefix[0])].size
		delete(trie.root.children, string(prefix[0]))
		return true
	}

	return false
}

func (trie *Trie) contains(elem string) *trienode {
	if len(elem) > 0 {
		tmp := trie.root
		for i := 0; i < len(elem); i++ {
			if _, ok := tmp.children[string(elem[i])]; !ok {
				return nil
			}
			tmp = tmp.children[string(elem[i])]
		}
		return &tmp
	}

	return nil
}

type trienode struct {
	isEndOfWord bool
	size        int
	children    map[string]trienode
}

func newTrieNode() trienode {
	return trienode{false, 0, make(map[string]trienode)}
}
