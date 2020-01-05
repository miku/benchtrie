package benchtrie

import (
	"log"
	"strings"
)

type Tree struct {
	Nodes map[string]*Tree
}

func NewTree() *Tree {
	return &Tree{
		Nodes: make(map[string]*Tree),
	}
}

func (tree *Tree) Print(indent string) {
	for k, t := range tree.Nodes {
		arrow := ""
		if t != nil {
			arrow = "↴"
		}
		log.Println(indent + k + arrow)
		if t != nil {
			t.Print(indent + "\t")
		}
	}
}

func (tree *Tree) Insert(list []string) {
	if len(list) == 0 {
		return
	}
	idx := len(list) - 1
	element := list[idx]

	node, ok := tree.Nodes[element]
	if !ok {
		if idx == 0 {
			tree.Nodes[element] = nil
		} else {
			t := NewTree()
			t.Insert(list[:idx])
			tree.Nodes[element] = t
		}
	} else {
		if node == nil {
			node = NewTree()
			tree.Nodes[element] = node
		}
		node.Insert(list[:idx])
	}
}

func (tree *Tree) LookupHost(hostname string) bool {
	parts := strings.Split(hostname, ".")
	var (
		p  = len(parts) - 1
		k  = tree
		ok bool
	)
	for {
		k, ok = k.Nodes[parts[p]]
		if !ok {
			return false
		}
		p--
		if p == -1 {
			return true
		}
	}
}
