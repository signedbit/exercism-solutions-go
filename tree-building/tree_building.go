package tree

import (
	"errors"
)

const RootID = 0

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func (n *Node) AddChild(c *Node) {
	n.Children = append(n.Children, c)
}

// Build has time & space O(N)+O(N) = O(2N) = O(N)
func Build(records []Record) (*Node, error) {
	size := len(records)
	if size == 0 {
		return nil, nil
	}

	nodes := make([]*Node, size)     // space: O(N)
	buckets := make([]*Record, size) // space: O(N)

	// create nodes
	// time: O(N)
	for _, r := range records {
		id := r.ID
		if id < 0 || id >= size {
			return nil, errors.New("ID is out of bounds")
		}
		if buckets[id] != nil {
			return nil, errors.New("duplicate node")
		}
		if id == r.Parent && id != RootID {
			return nil, errors.New("cannot be own parent")
		}
		if id < r.Parent {
			return nil, errors.New("invalid parent")
		}
		r := r
		buckets[id] = &r
		nodes[id] = &Node{ID: id}
	}

	// create pointers to children
	// time: O(N)
	for _, n := range nodes {
		if n.ID == RootID {
			continue
		}
		parentId := buckets[n.ID].Parent
		if nodes[parentId] == nil {
			return nil, errors.New("parent doesn't exist")
		}
		nodes[parentId].AddChild(n)
	}

	return nodes[RootID], nil
}
