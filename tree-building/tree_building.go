package tree

import (
	"errors"
	"sort"
)

// Record Identifies the ID and Parent
type Record struct {
	ID, Parent int
}

// Node defines the node structure
type Node struct {
	ID       int
	Children []*Node
}

// Build builds the tree.
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))
	for i, record := range records {
		if record.ID != i || record.Parent > record.ID || record.ID > 0 && record.Parent == record.ID {
			return nil, errors.New("invalid record")
		}
		// New node
		nodes[i] = &Node{ID: record.ID}

		if record.ID > 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
		}
	}

	return nodes[0], nil

}
