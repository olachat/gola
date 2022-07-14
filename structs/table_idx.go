package structs

import (
	"fmt"
	"sort"
)

// IdxNode defines a node(column) in a index
type IdxNode struct {
	Column
	ColName  string
	Children []*IdxNode
	maxOrder int
	Order    int
	Parent   *IdxNode
}

// GetChildren returns child *IdxNode with given colName
// If not found, a new child node will be created and returned
func (n *IdxNode) GetChildren(colName string) *IdxNode {
	for _, child := range n.Children {
		if child.ColName == colName {
			return child
		}
	}

	child := &IdxNode{
		ColName: colName,
		Parent:  n,
		Order:   n.GetNewOrder(),
	}
	n.Children = append(n.Children, child)

	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].ColName < n.Children[j].ColName
	})
	return child
}

// GetNewOrder returns a new order in the current *IdxNode tree
// Only root node will be handling this func
func (n *IdxNode) GetNewOrder() int {
	if n.Parent != nil {
		return n.Parent.GetNewOrder()
	}

	n.maxOrder++
	return n.maxOrder
}

// InterfaceName returns the interface name for current index node
func (n *IdxNode) InterfaceName() string {
	if len(n.Children) == 0 {
		return "orderReadQuery"
	}
	return fmt.Sprintf("iQuery%d", n.Order)
}

// StructName returns the sturct name for current index node
// The struct implements coresponding interface
func (n *IdxNode) StructName() string {
	return fmt.Sprintf("idxQuery%d", n.Order)
}

// GetAllChildren returns all children in current index node tree
// result is sorted in node's order ascending
func (n *IdxNode) GetAllChildren() []*IdxNode {
	result := n.Children
	for _, child := range n.Children {
		result = append(result, child.GetAllChildren()...)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Order < result[j].Order
	})
	return result
}

// String returns string representation of current index node tree
func (n *IdxNode) String(prefix string) string {
	s := prefix
	s += fmt.Sprintf("%s[%d]\n", n.ColName, n.Order)

	for _, c := range n.Children {
		s += c.String(prefix + "\t")
	}

	return s
}
