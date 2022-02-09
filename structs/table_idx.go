package structs

import (
	"fmt"
	"sort"
)

type idxNode struct {
	Column
	ColName  string
	Children []*idxNode
	maxOrder int
	Order    int
	Parent   *idxNode
}

func (n *idxNode) GetChildren(colName string) *idxNode {
	for _, child := range n.Children {
		if child.ColName == colName {
			return child
		}
	}

	child := &idxNode{
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

func (n *idxNode) GetNewOrder() int {
	if n.Parent != nil {
		return n.Parent.GetNewOrder()
	}

	n.maxOrder += 1
	return n.maxOrder
}

func (n *idxNode) InterfaceName() string {
	if len(n.Children) == 0 {
		return "orderReadQuery"
	}
	return fmt.Sprintf("idxQuery%d", n.Order)
}

func (n *idxNode) GetAllChildren() []*idxNode {
	result := n.Children
	for _, child := range n.Children {
		result = append(result, child.GetAllChildren()...)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Order < result[j].Order
	})
	return result
}

func (n *idxNode) String(prefix string) string {
	s := prefix
	s += fmt.Sprintf("%s[%d]\n", n.ColName, n.Order)

	for _, c := range n.Children {
		s += c.String(prefix + "\t")
	}

	return s
}
