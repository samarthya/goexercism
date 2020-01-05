package tree

import (
	"fmt"
	"sort"
)

// From the test cases.
type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

/**
 * AddToParent
 */
func addToParent(tree *Node, node *Node, parentID int) {

	added := false
	if tree != nil {
		for {
			if parentID != tree.ID {
				chil := tree.Children
				for _, e := range chil {
					if e.ID == parentID {
						e.Children = append(e.Children, node)
						added = true
						return
					}
				}

				if added == false {
					// The element parent could not be located.
				}
			}
		}
	}

	// Add the element as the first element itself.
	tree = node
}

func Build(records []Record) (tree *Node, err error) {

	var newNode *Node

	if records == nil || len(records) == 0 {
		return nil, nil
	}

	// First sort the element in the list so that the no parent is always the first.
	sort.Slice(records, func(i, j int) bool {
		return records[i].Parent < records[j].Parent
	})

	fmt.Printf(" The object : %v", records)

	for _, r := range records {
		if tree == nil {
			tree = &Node{r.ID, []*Node{}}
		} else {
			newNode = &Node{r.ID, []*Node{}}
			addToParent(tree, newNode, r.Parent)
		}
	}

	return tree, nil
}
