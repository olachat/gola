package structs

import "sort"

type idxNode struct {
	ColName  string
	Children []*idxNode
}

func (n *idxNode) GetChildren(colName string) *idxNode {
	for _, child := range n.Children {
		if child.ColName == colName {
			return child
		}
	}

	child := &idxNode{
		ColName: colName,
	}
	n.Children = append(n.Children, child)

	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].ColName < n.Children[j].ColName
	})
	return child
}

func (n *idxNode) String(prefix string) string {
	s := prefix
	s += n.ColName + "\n"

	for _, c := range n.Children {
		s += c.String(prefix + "\t")
	}

	return s
}
