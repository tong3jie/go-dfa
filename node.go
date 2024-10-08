package dfa

type Node struct {
	IsEnd bool
	Value string
	Child map[rune]*Node
}

func newNode(val string) *Node {
	return &Node{
		IsEnd: false,
		Value: val,
		Child: make(map[rune]*Node),
	}
}

type Trie struct {
	root *Node
	size int
}

func (t *Trie) Root() *Node {
	return t.root
}

func (t *Trie) Insert(key string) {
	curNode := t.root
	for _, v := range key {
		if curNode.Child[v] == nil {
			curNode.Child[v] = newNode(string(v))
		}
		curNode = curNode.Child[v]
		curNode.Value = string(v)
	}

	if !curNode.IsEnd {
		t.size++
		curNode.IsEnd = true
	}
}

func (t *Trie) Remove(key string) {
	words := []rune(key)
	curNode := t.root
	nodes := make([]*Node, 0, len(words))
	for _, v := range words {
		if curNode.Child[v] == nil {
			return // 键不存在，直接返回
		}
		curNode = curNode.Child[v]
		nodes = append(nodes, curNode)
	}

	if !curNode.IsEnd {
		return // 键不是完整的单词，直接返回
	}

	curNode.IsEnd = false
	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i].IsEnd || len(nodes[i].Child) > 0 {
			break
		}
		delete(nodes[i-1].Child, words[i])
	}
	t.size--
}

func (t *Trie) PrefixMatch(key string) []string {
	node, _ := t.findNode(key)
	if node == nil {
		return nil
	}
	return t.walk(node)
}

func (t *Trie) walk(node *Node) []string {
	var result []string
	t.walkHelper(node, "", &result)
	return result
}

func (t *Trie) walkHelper(node *Node, prefix string, result *[]string) {
	if node.IsEnd {
		*result = append(*result, prefix)
	}
	for char, child := range node.Child {
		t.walkHelper(child, prefix+string(char), result)
	}
}

func (t *Trie) findNode(key string) (node *Node, index int) {
	curNode := t.root
	f := false
	for k, v := range key {
		if f {
			index = k
			f = false
		}
		if curNode.Child[v] == nil {
			return nil, index
		}
		curNode = curNode.Child[v]
		if curNode.IsEnd {
			f = true
		}
	}

	if curNode.IsEnd {
		index = len(key)
	}

	return curNode, index
}

func (t *Trie) Child(key string) *Node {
	node, _ := t.findNode(key)
	return node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(""),
		size: 0,
	}
}
