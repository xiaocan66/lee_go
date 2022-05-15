package lee

import (
	"strings"
)

type node struct {
	pattern  string  // Route to be matched /p/:lang
	part     string  // part of the route /:lang
	children []*node // child node
	isWild   bool    // Whether to match exactly part contains '*' or ':'
}

//mathchChild Returns the first node that matches successfully.
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren Returns all matching nodes
func (n *node) mathchChildren(part string) []*node {
	var nodes []*node
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//insert define the method to insert a routing rule
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return

	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search Define a method to find out whether there is a corresponding routing rule through 'parts'
func (n *node) search(parts []string, height int) *node {
	//Directly truncate when ‘*’ appears
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if len(n.pattern) == 0 {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.mathchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
