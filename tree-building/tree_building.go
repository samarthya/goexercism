package tree

import (
	"sort"
	"errors"
)

// From the test cases.
type Record struct {
	ID, Parent int
}

// Node structure
type Node struct {
	ID       int
	Children []*Node
}


// A recursive approach to check if the new Node to add exists.
func containsTheNode(tree *Node, nodeId int) bool {
	if tree == nil {
		return false
	} 

	traverse := tree
	if traverse.ID == nodeId {
		return true
	}

	for _, val := range traverse.Children {
		return containsTheNode(val, nodeId)
	}

	return false
}

func addToTheParent(head *Node, newNode *Node, parent int) (*Node, error)  {
	if head != nil {
		if head.ID == parent {
			head.Children = append(head.Children, newNode)
			children := head.Children
			sort.Slice(children, func(i,j int) bool {
				return children[i].ID < children[j].ID
			})
			return head, nil
		} else {
			for _, val := range head.Children {
				if val.ID == parent {
					val.Children = append(val.Children, newNode)
					children := val.Children
					sort.Slice(children, func(i,j int) bool {
						return children[i].ID < children[j].ID
					})
					return head, nil
				}
			}
			// No parent found for this child, in the head.
			return nil, errors.New(" error. ")
		}
	} else {
		// The parent is not yet added.
		head = &Node { 
			parent,
			nil ,
		}

		return addToTheParent(head, newNode, parent)
	}
	return head, nil
}

// Build - Builds the tree.
func Build(records []Record) (tree *Node, err error) {
	// The records only contain an ID number and a parent ID number. 
	// var newNode *Node
	
	if records == nil {
		return nil, nil
	}
	
	// Sort the elements by index.
	sort.Slice(records, func(i,j int) bool {
		return records[i].ID < records[j].ID
	})

	// Iterate through the records and add one by one.
	for i, v := range records {

		if containsTheNode(tree, v.ID) {
			return nil, errors.New(" exists")
		}
		
		if i != v.ID {
			return nil, errors.New(" non continuous")
		}

		// Iterate through the records presented.
		if v.ID == v.Parent {
			// Root has the parent == id
			if tree == nil {
				// Root element found.
				tree = &Node{v.ID, nil, }
			}
		} else {

			if tree == nil {
				return nil, errors.New(" root not found ")
			}
			

			tree, _ = addToTheParent(tree, &Node{v.ID, nil,}, v.Parent)
		}
	}


	return tree, nil
}
