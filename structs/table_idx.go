package structs

import "sort"

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

func (n *idxNode) String(prefix string) string {
	s := prefix
	s += n.ColName + "\n"

	for _, c := range n.Children {
		s += c.String(prefix + "\t")
	}

	return s
}
